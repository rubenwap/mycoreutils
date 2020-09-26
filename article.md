# Upskill yourself by recreating GNU Coretools in Go

One of the most popular advice to gain skills programming is to reimplement tools that already exist. For instance, if you wanted to learn how to do a CRUD API, I am pretty sure you probably practiced with a "To do" app. The point of this is not to create something new and groundbreaking, but to use well known ideas to improve what you know. Other examples you might have seen are things such as implementing a chat server to learn about websockets, a weather app, a Hacker News reader...

Today I am talking about something I consider more interesting: Reimplementing the GNU Coretools utilities. I am not the first person doing this, but I am convinced that you are going to come out of this exercise with way better skills in your language of choice.

So what are the Coreutils? If you use Linux or Mac, you have probably used a bunch of them. Perhaps on Windows too in some cases. Take a look at this link for more information. 

You can approach this project in many ways. In my case, I know that the point of coreutils is that they are really small and really simple, they do just one task but they do it well. In the original repo in C, most of the tools are just one single file. I wanted to avoid copying the structure, so often I will 