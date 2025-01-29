const getInnerHTML = (id: string): string => {
    const elemnent = document.getElementById(id) as HTMLInputElement;
    return elemnent ? elemnent.value : "";
}