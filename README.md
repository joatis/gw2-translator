# gw2-translator
A project work translating runes in GuildWars 2 screen shots into English.

## Overview
A hybrid OCR pipeline. I used ImageMagick for thresholding and noise reduction, then orchestrated a Tesseract engine with a custom-trained language data set for New Krytan runes, all running in a containerized Go environment.

## The Stack:
1. Ollama - For a local LLM
2. Tesseract - Custom OCR
3. ImageMagick - For pre-processing images so the OCR can "see" better
4. Docker Desktop - Containerization for Go and PostGreSQL to save the translations.



