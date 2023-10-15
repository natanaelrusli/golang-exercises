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

### Is more worker good?

The optimal number of workers depends on various factors, including the nature of the tasks, the characteristics of the system, and the available resources. Having more workers can potentially increase parallelism and speed up the overall execution of tasks, but it's not always the case.

Here are some considerations:

- #### Task Nature:

  If tasks are CPU-bound, having more workers than the number of available CPU cores might not lead to better performance. On the other hand, for I/O-bound tasks (like downloading files), having more workers can be beneficial as they can overlap I/O operations.

- #### System Characteristics:

  The number of workers should be adjusted based on the system's capacity. If you have a system with a limited number of cores, having too many workers might lead to contention and degrade performance.

- #### Resource Constraints:

  Each worker consumes system resources, so there's a trade-off. Too many workers can lead to increased memory consumption and contention for resources.

- #### Network and Disk Speed:
  If the tasks involve network or disk operations, having more workers can help overlap latency, but there's a limit to the improvement.

It's often a good idea to experiment with different numbers of workers to find the optimal balance for your specific scenario. Consider monitoring system resources (CPU, memory, etc.) during execution to ensure that your application is not becoming resource-bound.

In the provided example, you can adjust the maxWorker value in the NewWorkerPool function to experiment with different numbers of workers. Keep in mind that setting the number of workers significantly higher than the number of available CPU cores might not always lead to better performance, especially for CPU-bound tasks.
