# Image Downloader

### Objectives

- Able to use goroutine
- Able to use channel

### Description

Write a program that downloads multiple images from a list of URLs concurrently, using a worker pool pattern. Use a channel to pass the URLs to the worker goroutines, and a separate channel to receive the downloaded images. Implement error handling and retry logic for failed downloads.

### Examples

Sample expected output (more or less)

```
Downloading image 1 from http://example.com/image1.jpg
Downloading image 2 from http://example.com/image2.jpg
Downloading image 3 from http://example.com/image3.jpg
Downloading image 4 from http://example.com/image4.jpg
Downloading image 5 from http://example.com/image5.jpg
Downloading image 6 from http://example.com/image6.jpg
Downloading image 7 from http://example.com/image7.jpg
Downloading image 8 from http://example.com/image8.jpg
Downloading image 9 from http://example.com/image9.jpg
Downloading image 10 from http://example.com/image10.jpg
Downloading image 11 from http://example.com/image11.jpg
Downloading image 12 from http://example.com/image12.jpg
Downloading image 13 from http://example.com/image13.jpg
Downloading image 14 from http://example.com/image14.jpg
Downloading image 15 from http://example.com/image15.jpg
Downloading image 16 from http://example.com/image16.jpg
Downloading image 17 from http://example.com/image17.jpg
Downloading image 18 from http://example.com/image18.jpg
Downloading image 19 from http://example.com/image19.jpg
Downloading image 20 from http://example.com/image20.jpg
Downloaded image 1 successfully
Downloaded image 2 successfully
Downloaded image 3 successfully
Downloaded image 4 successfully
Downloaded image 5 successfully
Downloaded image 6 successfully
Downloaded image 7 successfully
Downloaded image 8 successfully
Downloaded image 9 successfully
Downloaded image 10 successfully
Downloaded image 11 successfully
Downloaded image 12 successfully
Downloaded image 13 successfully
Downloaded image 14 successfully
Downloaded image 15 successfully
Downloaded image 16 successfully
Downloaded image 17 successfully
Downloaded image 18 successfully
Downloaded image 19 successfully
Downloaded image 20 successfully
```
