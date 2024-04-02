# TwitterSentiment
Does twitter sentiment analysis correlate to stock prices

## Requirements
- Python 3.9-3.11 (https://www.python.org/downloads/release/python-3118/)  
- pip requirements. Do:  
    ```pip install -r requirements.txt```
- If using a GPU: 
    - an NVIDIA graphics card
    - CUDA Toolkit 12.3 (https://developer.nvidia.com/cuda-12-3-0-download-archive)
    - cuDDN SDK 8.6.0 (https://developer.nvidia.com/rdp/cudnn-archive)

## Notes:
See notes in code. All of them are listed here for easy conversion to GitHub issues
- only using Apple data at the moment. unsure if its better to train per stock or all at once
- this scales the values for supposedly better allignment, but may not give intended results 
- reshaping needs to be done properly according to accurate LSTM inputs. see https://towardsdatascience.com/implementation-differences-in-lstm-layers-tensorflow-vs-pytorch-77a31d742f74
-  model can be changed. right now there's 3 layers of LSTM and Dropout however better results may be achived with different LSTM layering
- does not work due to improper shaping. see above link in NOTES