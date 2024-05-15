/*
 *  WordCount.java
 *
 *  Figure 17.8 with a few modifications by Lamont Samuels
 * 
 *  PLEASE NOTE THIS CODE WILL NOT COMPILE. This only a snippet of the solution.
 *
 * From "The Art of Multiprocessor Programming",
 * by Maurice Herlihy and Nir Shavit.
 * Copyright 2006 Elsevier Inc. All rights reserved.
 */

public 
public class WordCount {
    static List<String> text;
    static int numThreads = ...
    public static void main(String[] args) {
        text = readBook("document.tex");
        List<List<String> inputs = splitInputs(text, numThreads);
        MapReduce<List<String>>, String, Long, Long> mapReduce = new MapReduce(); 
        mapReduce.setMapperSupplier(WordCount.Mapper::new); 
        mapReduce.setReducerSupplier(WordCount.Reducer::new);
        mapReduce.setInput(inputs);
        Map<String, Long> map = mapReduce.call();
        displayOutput(map);
    }
}