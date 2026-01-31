+++
title = "InstanceListPattern"
date = 1350247781
category = "Old"
tags = ["Old"]
github = "https://github.com/alexanderbrevig/InstanceListPattern"
visibility = "public"
languages = "C#"
interest_score = 0
+++

## About

Easily make your object maintain an internal list of instances, relies on the programmer to Dispose();
## Original README

InstanceListPattern
===================

Easily make your object maintain an internal list of instances, relies on the programmer to Dispose();

You can install from Package Manager Console:

    PM> Install-Package InstanceListPattern


Here is an example:

    namespace InstanceListPatternSample
    {
        class Person : InstanceList<Person>
        {
            public Person(string name)
            {
                this.Name = name;
            }

            public string Name { get; set; }

            public override string ToString()
            {
                return Name;
            } 
        }
        class Program
        {
            static void Main(string[] args)
            {
                Person p1 = new Person("p1");
                Person p2 = new Person("p2");

                PrintAllPersons(); 
                
                Person p3 = new Person("p3");

                PrintAllPersons();

                p3.Dispose();

                PrintAllPersons();

                Console.Read();
            }

            private static void PrintAllPersons()
            {
                Console.WriteLine("All instances:");
                foreach (var p in Person.Instances) {
                    Console.WriteLine(p);
                }
            }
        }
    }

Print output:

    All instances:
    p1
    p2
    All instances:
    p1
    p2
    p3
    All instances:
    p1
    p2
