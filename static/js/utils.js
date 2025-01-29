"use strict";
const getInnerHTML = (id) => {
    const elemnent = document.getElementById(id);
    return elemnent ? elemnent.value : "";
};
