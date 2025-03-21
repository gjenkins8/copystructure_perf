#!/bin/bash

printf "\n\n===== Running Wasm build ======\n\n"
make -C wasm build run

printf "\n\n===== Running native build ====\n\n"
make -C native build run