#!/usr/bin/env sh

# Upload the directories.
(
  (
    echo "rm dexeq_out" &&
    echo "rm hot_tokens_api" &&
    echo "rm dexeq_out.zip" &&
    echo "rm restart.sh" &&
    echo "put ./scripts/output/dexeq_out.zip" &&
    echo "put -r ./chains" &&
    echo "put -r ./monitor" &&
    echo "put ./scripts/restart.sh"
  ) | sftp root@135.181.132.36:"/root/DexEqualizer/"
)

# Unpack.
ssh -t root@135.181.132.36 'cd ~/DexEqualizer/ && chmod +x restart.sh && unzip -o dexeq_out.zip'