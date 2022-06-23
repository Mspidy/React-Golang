
import axios from 'axios';
import { useEffect, useState } from 'react';
import './App.css';

function App() {
  const [gadgets, setGadget]= useState([])


  useEffect(()=>{
    async function getAllGadget(){
      try{
        const gadgets = await axios.get("http://localhost:8088/gadgets")
        console.log(gadgets.data)
        setGadget(gadgets.data)
      }catch (error){
        console.log(error)
      }
    }
    getAllGadget()
  }, [])
  return (
    <div className="App">
      <h1 style={{color:"red"}}>This is All Electronics Gadget's</h1>
      {
        gadgets.map((gadget, i)=>{
          return(
            <h2 key={i}>{gadget.id} {gadget.brand} {gadget.gadget} {gadget.price}</h2>
          )
        })
      }
    </div>
  );
}

export default App;
