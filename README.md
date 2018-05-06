# TIMEIN

Show your data by timezones.

![demo](https://github.com/ashlinchak/timein/blob/master/examples/timein.jpg)

### Usage
1. Compile this repository or make executable `bin/timein` file.
2. Prepare `data.json` and `config.json` files and leave them alongside executable `timein` file.
3. Run command (see the Examples section)

### Examples
Valid commands:
```
$ timein
$ timein berlin
$ timein berlin -t=21:00
```

Where:
* `berlin` - tag for filtering;
* `-t` - your local time. You want to see times on this time in other timezones that were filled in `data.json` file.

### Valid time zones
https://en.wikipedia.org/wiki/List_of_tz_database_time_zones