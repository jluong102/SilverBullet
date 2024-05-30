# Silver Bullet
Silver Bullet is an SRE tool desgined to to automaticlly 
run a remediation script an attempt to fix the issue without 
any user intervention.

## Setup
To run Silver Bullet you will first need to setup a general
YAML config file such as [example](./example/config.yaml) provided.
By default, this binary will look for `/etc/silverbullet/config.yaml`
but this can be changed using the `-c` option.

The YAML config file takes the following fields:
* [required]`bullets` -> List of strings of the "bullet" YAML files
to be used.
* [optional]`log` -> Directory to store logs.
* [optional] oor -> Directory used to mark "bullet" out of rotation.

### Config Example
```yaml
---
bullets:
  - examples/bullet.yaml
log: /var/silverbullet/log
oor: /etc/silverbullet/oor
```

Next, you will need to setup a [YAML "bullet" file](./example/bullet.yaml). 
Bullets will define what needs to be monitored and what actions 
should be taken when on a bad exit. 

Each bullet can be setup using the 
following fields:

* [required]`monitor` 
	* [required]`interval` -> Unsigned integer on how often the monitoring
check should be run.
	* [required]`script` -> The script (or binary) that should be run to check
if that status is okay.
	* [required]`good` -> A list of exitcodes that represents a successful check.
	* [required]`bad` -> A string/int list key/pair used to determine what remedy 
script (string) should be run on a bad exit (int).
* [required]`remedy`
	* [required]`(name_of_remedy)` -> For each bad exit script that is defined 
in monitoring. 
	* [required]`script` -> The script (or binary) that should be run to attempt to 
remedy the bad status found in the monitoring check.
	* [requried]`interval` -> An unsigned integer representing a wait period before 
checking the status again.
	* [required]`try` -> An unsigned integer representing how many times to run the 
remedy script before marking the bullet out-of-rotation. Using 0 will make infinite
attempts.

### Bullet Example
```yaml
---
monitor:
  interval: 240
  script: examples/check.sh
  good:
    - 0
  bad:
    example_fix:
      - 1
remedy:
  example_fix:
    script: examples/resolve.sh
    interval: 240
    try: 3
```
