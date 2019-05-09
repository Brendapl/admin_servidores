package main

import  "fmt"
import "math/rand"

//Vector que crea un Vector de longitud N
func vectorZero(tamaño int) []float64{

  vector := make([]float64, tamaño)

  return vector

}

///////////////////////////////////////////////////

//Vector con números aleatorios
func vectorRandom(tamaño int) []float64{

  vector := make([]float64, tamaño)
  for i := range vector {
    vector[i] = rand.Float64()
  }
  return vector
}

///////////////////////////////////////////////////

//Matriz cuadrada de zeros
func matrizZero(tamaño int) [][]int{

  matriz := make([][]int, tamaño)
  for i := range matriz {
      matriz[i] = make([]int, tamaño)
  }

  return matriz

}

///////////////////////////////////////////////////

//Matriz con números aleatorios
func matrizRandom(tamaño int) [][]int{

  matriz := matrizZero(tamaño)
  for i := range matriz {
    for j := range matriz{
      matriz[i][j] = rand.Intn(100)
    }
  }
  return matriz
}

///////////////////////////////////////////////////

//Multiplicación de Matriz con Vector
func matrizVector(tamaño int) []int{

  vector := make([]int, tamaño)
  for i := range vector {
    vector[i] = rand.Intn(100)
  }

  matriz := matrizZero(tamaño)
  for i := range matriz {
    for j := range matriz{
      matriz[i][j] = rand.Intn(100)
    }
  }
  multiplica := make([]int, tamaño)

  for i := range matriz{
    columna := 0
    fila := matriz[i]
    for j := range fila{
      columna += fila[j] * vector[j]
    }
    multiplica[i] = columna
  }
  fmt.Println("Vector", vector)
  fmt.Println("Matriz", matriz)
  fmt.Println("Resultado")
  return multiplica
}



///////////////////////////////////////////////////

func main(){

  //Hello World
  fmt.Println("Hola Mundo!")
  fmt.Println("---------------------------------")
  fmt.Println("Vector con Zeros")
  fmt.Println(vectorZero(15))
  fmt.Println("---------------------------------")
  fmt.Println("Vector con Números Random")
  fmt.Println(vectorRandom(5))
  fmt.Println("---------------------------------")
  fmt.Println("Matriz con Zeros")
  fmt.Println(matrizZero(8))
  fmt.Println("---------------------------------")
  fmt.Println("Matriz con Números Random")
  fmt.Println(matrizRandom(10))
  fmt.Println("---------------------------------")
  fmt.Println("Matriz por Vector")
  fmt.Println(matrizVector(4))

}
