#!/bin/bash

# Membuat folder foods, drinks, dan snacks
mkdir foods drinks snacks

# Membuat file menu.txt di dalam setiap folder
touch foods/menu.txt drinks/menu.txt snacks/menu.txt

# Menambahkan item ke file menu.txt di dalam folder foods
echo "nasi goreng" >> foods/menu.txt
echo "ayam goreng" >> foods/menu.txt
echo "bubur ayam" >> foods/menu.txt

# Menambahkan item ke file menu.txt di dalam folder drinks
echo "kopi susu" >> drinks/menu.txt
echo "susu oat" >> drinks/menu.txt
echo "es teh" >> drinks/menu.txt

# Mendownload data snack dari URL dan menyimpannya ke file snacks.json
curl -o snacks/snacks.json https://gist.githubusercontent.com/nadirbslmh/c84ee3527272e52a8e6257d46e627f91/raw/74593cc3fe61ca4ff36e11ed8e3fffcfb1d0c602/snacks.json

# Mengambil data snack dari file snacks.json dan menambahkannya ke file menu.txt di dalam folder snacks
jq -r '.snacks | .[]' snacks/snacks.json >> snacks/menu.txt