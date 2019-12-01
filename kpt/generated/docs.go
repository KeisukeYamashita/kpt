// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "mdtogo"; DO NOT EDIT.
package generated

var DescShort = `Display package descriptions`
var DescLong = `
Display package descriptions.

` + "`" + `desc` + "`" + ` reads package information in given DIRs and displays it in tabular format.
Input can be a list of package directories (defaults to the current directory if not specifed).
Any directory with a Kptfile is considered to be a package.

    kpt desc [DIR]... [flags]
`
var DescExamples = `
	# display description for package in current directory
	kpt desc
	
	# display description for packages in directories with 'prod-' prefix
	kpt desc prod-*`

var GetShort = `Fetch a package from a git repository`
var GetLong = `
Fetch a package from a git repository.

    kpt get REPO_URI[.git]/PKG_PATH[@VERSION] LOCAL_DEST_DIRECTORY [flags]

  REPO_URI:

    URI of a git repository containing 1 or more packages as subdirectories.
    In most cases the .git suffix should be specified to delimit the REPO_URI from the PKG_PATH,
    but this is not required for widely recognized repo prefixes.  If get cannot parse the repo
    for the directory and version, then it will print an error asking for '.git' to be specified
    as part of the argument.
    e.g. https://github.com/kubernetes/examples.git
    Specify - to read Resources from stdin and write to a LOCAL_DEST_DIRECTORY.

  PKG_PATH:

    Path to remote subdirectory containing Kubernetes Resource configuration files or directories.
    Defaults to the root directory.
    Uses '/' as the path separator (regardless of OS).
    e.g. staging/cockroachdb

  VERSION:

    A git tag, branch, ref or commit for the remote version of the package to fetch.
    Defaults to the repository master branch.
    e.g. @master

  LOCAL_DEST_DIRECTORY:

    The local directory to write the package to.
    e.g. ./my-cockroachdb-copy

    * If the directory does NOT exist: create the specified directory and write
      the package contents to it
    * If the directory DOES exist: create a NEW directory under the specified one,
      defaulting the name to the Base of REPO/PKG_PATH
    * If the directory DOES exist and already contains a directory with the same name
      of the one that would be created: fail

  --pattern string
  
    Pattern to use for writing files.  May contain the following formatting verbs
    %n: metadata.name, %s: metadata.namespace, %k: kind (default "%n_%k.yaml")
`
var GetExamples = `
	# fetch package cockroachdb from github.com/kubernetes/examples/staging/cockroachdb
	# creates directory ./cockroachdb/ containing the package contents
	kpt get https://github.com/kubernetes/examples.git/staging/cockroachdb@master ./
	
	# fetch a cockroachdb
	# if ./my-package doesn't exist, creates directory ./my-package/ containing the package contents
	kpt get https://github.com/kubernetes/examples.git/staging/cockroachdb@master ./my-package/
	
	# fetch package examples from github.com/kubernetes/examples
	# creates directory ./examples fetched from the provided commit
	kpt get https://github.com/kubernetes/examples.git/@8186bef8e5c0621bf80fa8106bd595aae8b62884 ./`

var InitShort = `Initialize suggested package meta for a local config directory`
var InitLong = `
Initialize suggested package meta for a local config directory.

Any directory containing Kubernetes Resource Configuration may be treated as
remote package without the existence of additional packaging metadata.

* Resource Configuration may be placed anywhere under DIR as *.yaml files.
* DIR may contain additional non-Resource Configuration files.
* DIR must be pushed to a git repo or repo subdirectory.

Init will augment an existing local directory with packaging metadata to help
with discovery.

Init will:

* Create a Kptfile with package name and metadata if it doesn't exist
* Create a Man.md for package documentation if it doesn't exist.


    kpt init DIR [flags]
    
  DIR:
    
    Defaults to '.'. Init fails if DIR does not exist

  --description string
  
    short description of the package. (default "sample description")
  
  --name string
  
    package name.  defaults to the directory base name.
  
  --tag strings
  
    list of tags for the package.
  
  --url string
  
    link to page with information about the package.
`
var InitExamples = `
	    # writes Kptfile package meta if not found
	    kpt init ./ --tag kpt.dev/app=cockroachdb --description "my cockroachdb implementation"`

var ManShort = `Format and display package documentation if it exists`
var ManLong = `
Format and display package documentation if it exists.    If package documentation is missing
from the package or 'man' is not installed, the command will fail.

    kpt man LOCAL_PKG_DIR [flags]

  LOCAL_PKG_DIR:

    local path to a package.
`
var ManExamples = `
	# display package documentation
	kpt man my-package/
	
	# display subpackage documentation
	kpt man my-package/sub-package/`

var UpdateShort = `Update a local package with changes from a remote source repo`
var UpdateLong = `
Update a local package with changes from a remote source repo.

    kpt update LOCAL_PKG_DIR[@VERSION] [flags]

  LOCAL_PKG_DIR:
  
    Local package to update.  Directory must exist and contain a Kptfile to be updated.

  VERSION:

  	A git tag, branch, ref or commit.  Specified after the local_package with @ -- pkg@version.
    Defaults the local package version that was last fetched.

	Version types:

    * branch: update the local contents to the tip of the remote branch
    * tag: update the local contents to the remote tag
    * commit: update the local contents to the remote commit

  --strategy:
  
    Controls how changes to the local package are handled.  Defaults to fast-forward.

    * resource-merge: perform a structural comparison of the original / updated Resources, and merge
	  the changes into the local package.  See ` + "`" + `kpt help apis merge3` + "`" + ` for details on merge.
    * fast-forward: fail without updating if the local package was modified since it was fetched.
    * alpha-git-patch: use 'git format-patch' and 'git am' to apply a patch of the
      changes between the source version and destination version.
      **REQUIRES THE LOCAL PACKAGE TO HAVE BEEN COMMITTED TO A LOCAL GIT REPO.**
    * force-delete-replace: THIS WILL WIPE ALL LOCAL CHANGES TO
      THE PACKAGE.  DELETE the local package at local_pkg_dir/ and replace it
      with the remote version.
          
  -r, --repo string
  
    Git repo url for updating contents.  Defaults to the repo the package was fetched from.

  --dry-run
  
    Print the 'alpha-git-patch' strategy patch rather than merging it.

#### Env Vars

  KPT_CACHE_DIR:
  
    Controls where to cache remote packages when fetching them to update local packages.
    Defaults to ~/.kpt/repos/
`
var UpdateExamples = `
	# update my-package-dir/
	kpt update my-package-dir/
	
	# update my-package-dir/ to match the v1.3 branch or tag
	kpt update my-package-dir/@v1.3
	
	# update applying a git patch
	git add my-package-dir/
	git commit -m "package updates"
	kpt update my-package-dir/@master --strategy alpha-git-patch`
