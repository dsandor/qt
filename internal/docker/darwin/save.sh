#!/bin/bash

docker save dsandor/qt:darwin | gzip -n > darwin.tar.gz
