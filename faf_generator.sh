#!/bin/env/bash

TEST_DIRS_PATH=$(pwd)/test_dirs

seq 100 | xargs -I {} mkdir -p $TEST_DIRS_PATH/dir_{}

for DIR in $(ls $TEST_DIRS_PATH)
do
  seq 10 | xargs -I {} mkdir -p $TEST_DIRS_PATH/$DIR/sub_dir_{}
done
