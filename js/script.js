// Copyright (c) 2023 Mohamad All rights reserved
//
// Created by: Mohamad
// Created on: Sep 2023
// This file contains the JS functions for index.html

"use strict"

/*
* This function calculates the area of a triangle
*/
function calculate () {

  // Debugging statements
  console.log("The button was clicked")

  // Assign the user input to a variable
  const sideLength = parseInt (document.getElementById("side-length").value)

  // Calculate the apothem length
  const apothemLength = sideLength / (2 * Math.tan(Math.PI / 20));

  // Calculate the area of the icosagon
  const area = (20 * sideLength * apothemLength) / 2;

  // Output the area of an icosagon
  document.getElementById("answer").innerHTML = "Area of the icsoagon is " + area.toFixed(2) + " cm²"
}
