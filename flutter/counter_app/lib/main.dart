import 'package:flutter/widgets.dart';
import 'package:flutter/material.dart';

void main () {
  runApp(MaterialApp(
    theme: ThemeData(
      primarySwatch: Colors.purple
    ),
    home: CountHome(),
  ));
}

class CountHome extends StatefulWidget {
  @override
  _CountHomeState createState() => _CountHomeState();
}

class _CountHomeState extends State<CountHome> {
  var count = 0;

  void increment() {
    count++;
    setState(() {
      
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Contador app"),
      ),
      body: Center(
        child: Text("Contador\n$count", textAlign: TextAlign.center,)
      ),
      floatingActionButton: FloatingActionButton(
        child: Icon(Icons.add), 
        onPressed: () {
          increment();
        },
      ),
    );
  }
}