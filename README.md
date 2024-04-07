# TwitterSentiment
Does twitter sentiment analysis correlate to stock prices

## Requirements
- Python 3.9-3.11 (https://www.python.org/downloads/release/python-3118/)  
- pip requirements. Do:  
    ```pip install -r requirements.txt```
- If using a GPU: 
    - an NVIDIA graphics card
    - TensorFlow<2.11
    - CUDA Toolkit 11.2 (https://developer.nvidia.com/cuda-11.2.0-download-archive)
    - cuDNN SDK 8.1.0 (https://developer.nvidia.com/rdp/cudnn-archive) 
    - See link for installing cuDNN (https://docs.nvidia.com/deeplearning/cudnn/archives/cudnn-896/install-guide/)

## Notes:
See notes in code. All of them are listed here for easy conversion to GitHub issues
- only using Apple data at the moment. unsure if its better to train per stock or all at once
- this scales the values for supposedly better allignment, but may not give intended results 
- reshaping needs to be done properly according to accurate LSTM inputs. see https://towardsdatascience.com/implementation-differences-in-lstm-layers-tensorflow-vs-pytorch-77a31d742f74
-  model can be changed. right now there's 3 layers of LSTM and Dropout however better results may be achived with different LSTM layering
- does not work due to improper shaping. see above link in NOTES