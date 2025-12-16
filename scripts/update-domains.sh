#!/bin/bash
set -e

echo "Updating domain lists..."

# download disposable domains
echo "Downloading disposable domains..."
curl -s https://raw.githubusercontent.com/disposable-email-domains/disposable-email-domains/master/disposable_email_blocklist.conf -o data/disposable_domains.txt
echo "✓ Disposable domains: $(wc -l < data/disposable_domains.txt) domains"

# download free email domains
echo "Downloading free email domains..."
curl -s https://raw.githubusercontent.com/willwhite/freemail/master/data/free.txt -o /tmp/freemail_raw.txt

# remove domains that are in disposable list from free list
echo "Filtering free domains (removing disposable overlaps)..."
comm -23 <(sort /tmp/freemail_raw.txt) <(sort data/disposable_domains.txt) > data/free_domains.txt

# add header comment
echo "# Free Email Providers" > /tmp/free_header.txt
echo "# Source: https://github.com/willwhite/freemail" >> /tmp/free_header.txt
echo "# Automatically filtered to exclude disposable domains" >> /tmp/free_header.txt
cat /tmp/free_header.txt data/free_domains.txt > /tmp/free_final.txt
mv /tmp/free_final.txt data/free_domains.txt

echo "✓ Free domains: $(grep -v '^#' data/free_domains.txt | grep -v '^$' | wc -l) domains"

# cleanup
rm -f /tmp/freemail_raw.txt /tmp/free_header.txt

echo "Done! Lists updated successfully."
