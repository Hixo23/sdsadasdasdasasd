import axios from 'axios';
import {fetchData, data } from './fetchData'

export const deleteLink = (ID) => {
    axios.delete(`http://localhost:3000/links/${ID}`).then(res => {
    })
    setTimeout(() => {
        fetchData();
    }, 1000)
}
