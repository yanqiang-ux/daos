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
from general_utils import write_remote_file


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
        self.principal = principal

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
        # Add correct format
        tmp_principal = self.principal
        if "@" not in self.principal:
            tmp_principal = self.principal + "@"

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
        return "ACL Entry: {}".format(self.__str__())

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

    def str_to_entry(self, stdout):
        """Import ACL information from a provided ACL stdout string.

        Args:
            stdout (str): string containing ACL information to be imported.
        """
        for key, entry in enumerate(stdout.splitlines()):
            if not entry.startswith("#"):
                e_type, flags, princ, perm = entry.split(":")
            self.entries[key] = ACE(e_type, flags, princ, set(perm))

    def add(self, key, entry):
        """Add entry to ACL.

        Args:
            key (str): identifier for ACE to be added.
            entry (ACE object): new access control entry.

        Returns:
            dict: dictionary containing entries in ACL.

        """
        if key in self.entries:

            self.entries[key].add(entry)
        else:
            self.entries[key] = set(self.entries)

    def remove(self, key):
        """Remove an entry from ACL.

        Args:
            key (str): identifier for ACE to be removed.

        """
        if key in self.entries:
            for

    def overwrite(self, key, entry):
        """Overwrtie an existing ACL entry with a new entry.

        Args:
            key (str): key to identify which entry will be overwritten.
            entry (ACE object): new entry object.

        Raises:
            KeyError:
        """
        if key in self.entries:
            self.entries[key] = entry
        else:
            raise KeyError

    def update(self, key, field, value):
        """Update ACL with provided entries.

        Args:
            key (str): key to identify which entry will be updated.
            field (str): field within an entry to update.
                i.e. type, flags, principal, permissions
            value (str): value used to update provided field.
        """

    def delete(self):
        """Reset ACL object and clear all entries."""
        self.entries.clear()

    def update_file(self, hosts=None):
        """Update an ACL file from the entries.

        Each entry should be on a separate new line to comply with ACL format.

        Args:
            hosts (str, optional): list of hosts, if not hosts provided,
                file will be updated locally.
        """
        return write_remote_file(
            hosts, self.file, "\n".join(self.entries.values()))
