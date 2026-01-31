+++
title = "PicoFSM"
date = 1350081589
category = "Old"
tags = ["Old"]
github = "https://github.com/alexanderbrevig/PicoFSM"
visibility = "public"
languages = "C#"
interest_score = 0
+++

## About

Extremely lightweight FSM library, supports edge definitions in the form of TransitionToWhen(state,condition)
## Original README

PicoFSM
=======

Extremely lightweight FSM library, supports edge definitions in the form of TransitionToWhen(state,condition)

Install from Package Manager Console:

    PM> Install-Package picoFSM

Example
-------

    Machine machine = new Machine();

    var second = new State(
        "second",
        (s) => { Console.WriteLine("second.Update until TransitionTo"); return false; }
    );

    var first = new State(
        "first",
        (s) => { Console.WriteLine("first.Enter"); return true; },
        (s) => { Console.WriteLine("first.Update"); return true; },
        (s) => { Console.WriteLine("first.Exit"); machine.TransitionTo(second);  return true; }
    );

    bool trueOnce = true;
    /// This is a rule that forces first to be selected as the first state
    /// We could do machine.TransitionTo as part of this setup, 
    /// but this demonstrates the use of TransitionToWhen
    machine.TransitionToWhen(first, (s) => {
        bool tmp = trueOnce;
        trueOnce = false;
        return tmp;
    });

    machine.Update();
    machine.Update();
    machine.Update();
    machine.Update();
    machine.Update();
    Console.WriteLine("...");

Prints:

    first.Enter
    first.Update
    first.Exit
    second.Update until TransitionTo
    second.Update until TransitionTo
    second.Update until TransitionTo
    ...

Refer to another example here:
https://github.com/AlexanderBrevig/PicoFSM/blob/master/PicoFSMSample/Program.cs
