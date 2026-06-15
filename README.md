# gw2-translator
A project work translating runes in GuildWars 2 screen shots into English.

## Overview
A hybrid OCR pipeline. I used ImageMagick for thresholding and noise reduction, then orchestrated a Tesseract engine with a custom-trained language data set for New Krytan runes, all running in a containerized Go environment.

## The Stack:
1. Ollama - For a local LLM
2. Tesseract - Custom OCR
3. ImageMagick - For pre-processing images so the OCR can "see" better
4. Docker Desktop - Containerization for Go and PostGreSQL to save the translations.

### Setting up and Training Tesseract
1. In ordewr to translate New Krytan we need to get the font file from [Proper Dave's New Krytan Typeface](https://www.properdave.com/gw2/)
2. Download the file.
3. Create a .fonts folder in the container, and copy the ttf file into it:
   `mkdir -p ./fonts`
   Note that this creates it in the vscode user's path eg. "/home/vscode"
   `mv NewKrytan.ttf ~/.fonts/NewKrytan.ttf`
4. Refresh font cache:
   `fc-cache -fv`
5. Ensure that the new font is listed:
   `fc-list : family | grep -i "krytan"`
6. Clone the tesseract repo into the project root directory.
    `git clone https://github.com/tesseract-ocr/tesstrain.git`
7. Before I could train with tesstrain I needed to install pip, and create a venv 
  1. sudo apt update
  2. sudo apt install python3-pip
  3. sudo apt install python3.11-venv
  4. python3 -m venv ./tesstrain_env
8. Install tesstrains requirements:
   `tesstrain_env/bin/pip install -r requirements.txt`
9. Fetch the raw language configurations required by the engine:
   `make tesseract-langdata`
10. Within tesstrain/data/ create a folder structure to contain Ground truth text. ('qgw' is an unallocated ISO prefix. The qaa - qtz block used for constructed languages. I'm uding gw for GuildWars)
   `mkdir -p data/qgw_krytan-ground-truth`
11. Copy `/training_data/translator_training_data.txt` to `tesstrain/data/qgw_krytan-ground-truth/training_text.txt`.
    Then use the helper script to split that file into one-line `.gt.txt` files and generate matching line images and box files.
12. Generate the line file set from `tesstrain/`:
    ```
    ./tesstrain_env/bin/python3 generate_training_line_files.py \
      --text data/qgw_krytan-ground-truth/training_text.txt \
      --out-dir data/qgw_krytan-ground-truth \
      --font NewKrytan \
      --fonts-dir /home/vscode/.fonts \
      --fontconfig-tmpdir /tmp/text2image-fontconf \
      --ptsize 24 \
      --resolution 300
    ```
    This creates `line_0001.gt.txt`, `line_0001.tif`, `line_0001.box`, `line_0002.gt.txt`, etc.
13. Remove any stale `training_text.*` artifacts from `data/qgw_krytan-ground-truth` before running training. Only the generated `line_*.gt.txt`, `line_*.tif`, and `line_*.box` files should remain in that folder.
    ```bash
    rm -f data/qgw_krytan-ground-truth/training_text.{tif,gt.txt,box,lstmf}
    rm -f data/eng/qgw_krytan.*
    ```
14. Run the automated training command after the ground truth directory contains one-line image/text pairs.

    If you are fine-tuning from `START_MODEL=eng`, `TESSDATA` must point to a trainable (double-based) `eng.traineddata`, not the distro `tessdata` fast integer model. If your system tessdata is fast, download the `eng.traineddata` from `tessdata_best` and use that directory.

    ```bash
    make TESSDATA=../training_data \
      PY_CMD=./tesstrain_env/bin/python3 \
      training \
      MODEL_NAME=qgw_krytan \
      START_MODEL=eng \
      FONT_NAME="NewKrytan" \
      MAX_ITERATIONS=400
    ```
### NOTE: Had the idea to package up all the training data into the training data directory and copy artifacts to tesstrain.




