const express = require('express')
const app = express()
const port = 3000

app.use(express.json())
app.use(express.urlencoded({extended:true}))

app.get("/",(req,res)=>{
    res.status(200).send("Welcome to server local")
})

app.post("/post",(req,res)=>{
    let myJSON=req.body
    req.status(200).send(myJSON)
})

app.post("/postform",(req,res)=>{
    res.status(200).send(JSON.stringify(req.body))
})

app.listen(port,()=>{
    console.log(`Example app listening at http://localhost:${port}`)
})