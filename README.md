<!--
SPDX-FileCopyrightText: 2023 Dinesh Ravi

SPDX-License-Identifier: GPL-3.0-only
-->

# blackduckAPI

This api focuses on performing bulk snippet match to a given component and component version with path matching to the given filter word and given no, of entries

## Docker image
```
docker pull dineshr93/hubsm:1.0
```

## command

```
docker run dineshr93/hubsm:1.0 1.Hub_api_url 2.hub_token 3.projectversionlink 4.codelocation 5.matchToComponentVersion 6.matchString 7.EntriesToResolve
```

## alias image
```
alias hubsm='docker pull dineshr93/hubsm:1.0'
```

### example
```
docker run dineshr93/hubsm:1.0 https://blackduck.com/api/ Zzc1ZjliMmQtMWZmZi00NDBjLTliY2UtZGM0ODNkOTNmMYzNTlmZjIxLTMwZjMtNGFkZS1iMDRhLWFjMDBmMmU4MmUyNJ https://blackduck.com/api/projects/10470624-5c56-4a74-b145-98a4963f4405/versions/cfda0cdc-831a-4fd1-a200-ac3d9af4d14f https://blackduck.com/api/codelocations/e2344cc7-2c84-4432-9feb-d2cf62d3184f https://blackduck.com/api/components/55697cd6-efbf-418d-b231-96df78f46dd8/versions/62fb441c-300a-4515-9c8c-b9e9c655f422 PinyinIME 1000

hubsm https://blackduck.com/api/ Zzc1ZjliMmQtMWZmZi00NDBjLTliY2UtZGM0ODNkOTNmMYzNTlmZjIxLTMwZjMtNGFkZS1iMDRhLWFjMDBmMmU4MmUyNJ https://blackduck.com/api/projects/10470624-5c56-4a74-b145-98a4963f4405/versions/cfda0cdc-831a-4fd1-a200-ac3d9af4d14f https://blackduck.com/api/codelocations/e2344cc7-2c84-4432-9feb-d2cf62d3184f https://blackduck.com/api/components/55697cd6-efbf-418d-b231-96df78f46dd8/versions/62fb441c-300a-4515-9c8c-b9e9c655f422 PinyinIME 1000


```

# usage
```
hubsm "1.https://blackduck.com/api/" "2.TOKEN" 3.project_version_url 4.signature_scan_codelocation_url 5.matchto_components_version_url 6.match_string_or_path 7.entries_to_match

example:
project_version_url:
https://blackduck.com/api/projects/250a7fcb-b49d-4cda-9302-c888ad562024/versions/81df0df9-aa6d-4750-a1e7-6a049c6b4965

signature_scan_codelocation_url
https://blackduck.com/api/codelocations/3cf0a11a-e453-4d64-ba3b-5ee87aaaeabt

matchto_components_version_url
https://blackduck.com/api/components/7eac8f37-d9e5-4344-83d0-be0e9fd42a6a/versions/9b57e540-34c1-4f6c-89cd-a0137e1959u6
```
