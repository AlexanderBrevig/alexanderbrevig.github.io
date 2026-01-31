+++
title = "SuperSimplePubSub"
date = 1410282966
category = "Old"
tags = ["Old"]
github = "https://github.com/alexanderbrevig/SuperSimplePubSub"
visibility = "public"
languages = "C#"
interest_score = 0
+++

## About

C# SuperSimple PubSub implementation
## Original README

SuperSimplePubSub
=================

C# SuperSimple PubSub implementation

This is simple:

    class PubSubTest
    {
        enum Test { ONE, TWO };

        static void Main(string[] args)
        {
            string test = " world";
            Action<object> hi = (ob) => { Console.WriteLine("hi" + ob); };
            Pub.Sub["hello"] = hi;
            Pub.Sub["hello"] = (ob) => { Console.WriteLine("hello" + ob); };
            Pub.Sub["hello"](test);

            Pub.Sub[Test.ONE] = (ob) => { Console.WriteLine("enum" + ob); };
            Pub.Sub[Test.ONE](test);

            Pub.Sub.Remove("hello", hi);
            Pub.Sub["hello"](test);
            Console.ReadKey();
        }
    }
    
    //prints:
    //hi world
    //hello world
    //enum world
    //hello world
    
Be sure to:

    PM> Install-Package SuperSimplePubSub
