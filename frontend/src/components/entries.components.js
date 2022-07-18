import React , { useState, useEffect, setRefreshData  } from 'react';
import axios from "axios";
import { Button, Form , Container, Modal } from 'bootstrap';
import Entry from './single-entry.component';

const Entriess=() =>{

    const [entriers, setEntries]  = useState([])
    const [refreshData, setRefreshData] =  useState(false)
    const[ChangeEntry,setChangeEntry] =  useState({"change":false,"id":0})
    const[ChangeIngredient,setChangeIngredient] = useState({"change":false, "id":0})
    const[newIngredientName , setnewIngredientName] = useState("")
    const[addNewEntry, setAddNewEntry] =  useState({
        "dish":"",
        "ingredients":"",
        "calories":0, fat:0 })
        //InnitialLoadUp
        useEffect(()=>{
            getAllEntries();
        })
        //Once we Refresh we Get all Entries
        if(refreshData){
            setRefreshData(false);
            getAllEntries();
        }
    return(
        <div>
            <Container>
                <button onClick={()=> setAddNewEntry(true)}>Track Today's</button>
            </Container>
            <Container>
                {entriers!=null && entriers.map((entry, i)=>
                    (
                        <Entry  entryData={entry} deleteSingleEntry={deleteSingleEntry}
                                setChangeIngredient={setChangeIngredient} setChangeEntry={setChangeEntry}
                        />

                    )
                )}
            </Container>
        </div>
    );
    
}

function addSingleEntry(){
    setAddNewEntry(false)
    //Connecting to Our Backend
    var url = "https://localhost:8080/entry/create"
    //Mapping our data using axios
    axios.post(url,{
        "ingredients":newEntry.ingredients,
        "dish": newEntry.dish,
        "calories":newEntry.calories,
        "fat":parseFloat(newEntry.fat)
    }).then(response =>{
        if(response.status
            ==200){
                setRefreshData(true)
            }
    })

}

function deleteSingleEntry(id){
        var url = "https://localhost:8080/entry/delete"+id
        axios.delete(url,{
            
        }).then(response =>{
            if(response.status
                ==200){
                    setRefreshData(true)
                }
        })

}

function getAllEntries(){
    //Our Go Server
    var url = "https://localhost:8080/entry/create"
    axios.get(url, {responseType:'json'}).then(response =>{
        if(response.status == 200){
            setEntries(response.data)
        }
    })
}