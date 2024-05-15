/*
 *  mapreduce.java
 *
 *  Figure 17.8 with a few modifications by Lamont Samuels
 * 
 *  PLEASE NOTE THIS CODE WILL NOT COMPILE. This only a snippet of the solution.
 * 
 * From "The Art of Multiprocessor Programming",
 * by Maurice Herlihy and Nir Shavit.
 * Copyright 2006 Elsevier Inc. All rights reserved.
 */

public class MapReduce<IN, K, V, OUT> implements Callable<Map<K, OUT>> {

private List<IN> inputList;
private Supplier<Mapper<IN, K, V>> mapperSupplier;
private Supplier<Reducer<K, V, OUT>> reducerSupplier;
private static ForkJoinPool pool;

public MapReduce () {
    pool = new ForkJoinPool ();
    mapperSupplier = ... // Think of this supplier as entity supplies threads that represent Mappers
    reducerSupplier = ...  // Think of this supplier as entity supplies threads that represent Reducers 
}
public Map<K, OUT> call() {
    Set<Mapper<IN, K, V> mappers = new HashSet<();
    for (IN input : inputList) {
        Mapper<IN, K, V> mapper = mapperSupplier.get ();
        mapper.setInput(input);
        pool.execute(mapper);
        mappers.add(mapper);
    }
    Map<K, List<V> mapResults = new HashMap<>();
    for (Mapper<IN, K, V> mapper : mappers) {
        Map<K, V> map = mapper.join();
        for (K key : map.keySet ()) {
            mapResults.putIfAbsent (key, new LinkedList<>());
            mapResults.get(key).add (map.get (key)) ;
        }
    }
    Map<K, Reducer<K, V, OUT>> reducers = new HashMap<>(); 
    mapResults.forEach(
        (k, v) -> {
        Reducers <K, V, OUT> reducer = reducerSupplier.get();
        reducer.setInput (k, v);
        pool.execute(reducer);
        reducers.put(k, reducer);

    );
    Map<K, OUT> result = new HashMap«>();;
    reducers.forEach (
        (key, reducer) -> {
            result.put(key, reducer.join());
        }
    return result; 
}