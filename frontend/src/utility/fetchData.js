import axios from "axios";
import { ref } from 'vue'
export const data = ref(null);
export const fetchData = () => {

    axios.get('http://localhost:3000/links')
        .then(function (response) {
            data.value = response.data.links
        })
        .catch(function (error) {
            // handle error
            console.log(error);
        })
}

fetchData()

