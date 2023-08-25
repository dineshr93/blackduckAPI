<!--
SPDX-FileCopyrightText: 2023 Dinesh Ravi

SPDX-License-Identifier: GPL-3.0-only
-->

# blackduckAPI

This api focuses on performing bulk snippet match to a given component and component version with path matching to the given filter word and given no, of entries

# usage
```
hubsm https://blackduck.com/api/ "TOKEN" project_version_url signature_scan_codelocation_url matchto_components_version_url match_string_or_path entries_to_match

example:
project_version_url:
https://blackduck.com/api/projects/250a7fcb-b49d-4cda-9302-c888ad562024/versions/81df0df9-aa6d-4750-a1e7-6a049c6b4965

signature_scan_codelocation_url
https://blackduck.com/api/codelocations/3cf0a11a-e453-4d64-ba3b-5ee87aaaeabt

matchto_components_version_url
https://blackduck.com/api/components/7eac8f37-d9e5-4344-83d0-be0e9fd42a6a/versions/9b57e540-34c1-4f6c-89cd-a0137e1959u6
```
