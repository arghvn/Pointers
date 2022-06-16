smart pointer:
  Pointers have two problems, one is that if they are not erased after use, those values inside the pointer remain in memory, which is called memory leak.
Secondly, we may have several pointers pointing to the same place. Here, too, if we delete one of the pointers by mistake, the other pointers point to the wrong place, which also causes It becomes a bug.


The advantage of smart pointers over regular pointers is that the memory that such pointers point to is freed up quite intelligently, and we no longer need to free up them manually.
In fact, using such pointers makes memory management much easier and gives you various settings for memory management.

we use it later