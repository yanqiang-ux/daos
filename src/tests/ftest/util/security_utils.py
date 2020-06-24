#!/usr/bin/python
"""
  (C) Copyright 2020 Intel Corporation.

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

  GOVERNMENT LICENSE RIGHTS-OPEN SOURCE SOFTWARE
  The Government's rights to use, modify, reproduce, release, perform, display,
  or disclose this software are subject to the terms of the Apache License as
  provided in Contract No. 8F-30005.
  Any reproduction of computer software, computer software documentation, or
  portions thereof marked with this legend must also reproduce the markings.
"""

import os
from logging import getLogger
from logger_utils import TestLogger
from general_utils import write_file


class ACE(object):
    """A class for an ACL entry."""

    PERMISSIONS = {
        "read": "r",
        "write": "w",
        "create": "c",
        "delete": "d",
        "get-prop": "t",
        "set-prop": "T",
        "get-acl": "a",
        "set-acl": "A",
        "set-owner": "o",
    }

    def __init__(self, acl_type=None, flags=None, principal=None,
                 permissions=None):
        """ Create the ACE object.

        Args:
            acl_type (str): type of entry.
            flags (str): information on how the ACE should be interpreted.
            principal (str): the identity, specified in name@domain format.
            permissions (set, optional): type of user access to resources.
        """
        self.acl_type = acl_type.upper()
        self.flags = flags.upper()
        self.principal = principal if "@" in principal else principal + "@"

        if permissions and not isinstance(permissions, set):
            self.permissions = set(permissions)
        else:
            self.permissions = set([])

    def __str__(self):
        """Convert this ACE into a string of the following format:

            TYPE:FLAGS:PRINCIPAL:PERMISSIONS

        Returns:
            str: the string version of the ACE values.

        """
        return "{}:{}:{}:{}".format(
            self.acl_type,
            self.flags,
            tmp_principal,
            "".join(self.permissions))

    def __repr__(self):
        """Convert this ACE into a string.

        Returns:
            str: the representation string for the ACE.

        """
        return "ACL Entry: \n{}".format(self.__str__())

    def update(self, field, value):
        """Update entries field with provided value.

        Args:
            field (str): field within an entry to update.
                i.e. type, flags, principal, permissions
            value (str): value used to update provided field.
        """
        if field in self.__dict__:
            if field == "permissions":
                if not isinstance(value, set):
                    self.permissions = set(value)
                else:
                    self.permissions = value
            else:
                setattr(self, field, value)
        else:
            raise KeyError("No entry field named: {}".format(field))

    def grant_permission(self, permission):
        """Grant permissions to this entry.

        Args:
            permission (str): permission to be granted i.e. read, write, create

        Returns:
            list: list of permissions given to this entry

        """
        if permission in self.PERMISSIONS:
            self.permissions.add(self.PERMISSIONS[permission])
        return self.permissions

    def revoke_permission(self, permission):
        """Revoke permissions to this entry.

        Args:
            permission (str): permission to be revoked i.e. read, write, create

        Returns:
            list: list of permissions given to this entry.

        """
        if permission in self.PERMISSIONS and \
           self.PERMISSIONS[permission] in self.permissions:
            self.permissions.remove(self.PERMISSIONS[permission])

        return self.permissions


class ACL(object):
    """A class for managing an ACL."""

    FILE_EXT = "txt"

    def __init__(self, path="/tmp", filename="acl_file"):
        """Create a Access Control List object."""
        self.entries = {}
        self.file = os.path.join(path, filename)

    def __str__(self):
        """Get ACL entries into ACL file format.

        Returns:
            str: ACL file formatting, with each entry in a new line.

        """
        return "{}".format("\n".join(self.entries.values()))

    def entry_import(self, stdout):
        """Import ACL information from a provided ACL stdout string.

        Args:
            stdout (str): string containing ACL information to be imported.
        """
        for entry in stdout.splitlines():
            if not entry.startswith("#"):
                e_type, flags, princ, perm = entry.split(":")
                self.entries[princ] = ACE(e_type, flags, princ, perm)

    def delete(self, key):
        """Remove an entry from ACL.

        Args:
            key (str): identifier for ACE to be removed.

        Raises:
            KeyError: if key is not found in ACL entries.
        """
        if key in self.entries:
            self.entries.pop(key)
        else:
            raise KeyError("No entry found with '{}' key".format(key))

    def overwrite(self, acl_info):
        """Overwrtie an existing ACL entry with a new entry.

        Args:
            acl_info (str): acl file information that will be used to ovewrite
                current acl entries.
        """
        for line in acl_info:


    def update(self, key, entry):
        """Update ACL with provided entries.

        Args:
            key (str): key to identify which entry will be updated.
            entry (ACE object):
        """
        self.entries[key] = entry

    def clear(self):
        """Clear all ACL entries."""
        self.entries.clear()

    def update_file(self, hosts=None):
        """Update an ACL file from the entries.

        Each entry should be on a separate new line to comply with ACL format.

        Args:
            hosts (str, optional): list of hosts, if not hosts provided,
                file will be updated locally.
        """
        return write_file(hosts, self.file, "\n".join(self.entries.values()))
