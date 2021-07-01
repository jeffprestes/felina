async function getRentalContractData() {
  let contractNumberField = document.getElementById("contractNumber");
  const userInput = contractNumberField.value * 1;
  alert("valorInformadoPeloUsuario: " + userInput);
  try {
    const arrayRentalData = await smartContract.rentals(userInput);
    console.log(arrayRentalData);
    //Modo mais verboso
    let showLocator = document.getElementById("locator");
    showLocator.innerText = arrayRentalData[0];
    //Modo sintetico onde se concatena um comando junto a outro na mesma linha
    document.getElementById("renter").innerText = arrayRentalData[1];
    document.getElementById("addressHome").innerText = arrayRentalData[2];
    document.getElementById("rentalValue").innerText = arrayRentalData[3];
  } catch (err) {
    console.error(err);
    document.getElementById("locator").innerText = "";
    document.getElementById("renter").innerText = "";
    document.getElementById("addressHome").innerText = "";
    document.getElementById("rentalValue").innerText = "";
    contractNumberField.value = 0;
    alert("Houve um erro ao buscar o contrato de numero: " + userInput);
  }
}

async function autoLoadOwner() {
  try {
    const contractOwner = await smartContract.owner();
    console.log(contractOwner);
    document.getElementById("spanOwner").innerText = contractOwner;
  } catch (err) {
    console.error(err);
    alert("Houve um erro ao buscar o proprietário do contrato");
  }
}

async function saveFormData() {
  try {
    var tx;
    var txReceipt;
    tx = await smartContractWithSigner.registerRental(
      document.frmImovel.paramLocator.value,
      document.frmImovel.paramRenter.value,
      document.frmImovel.paramAddressHome.value,
      document.frmImovel.paramRentalValue.value
    );
    console.log("transacao enviada ao metamask. pendente...", tx);
    alert("Transação enviada... " + tx.hash + " aguarde a confirmação da Blockcnain...");
    txReceipt = await tx.wait();
    console.log("transacao processada...", txReceipt);
    if (txReceipt.status == 1) {
      alert("Transação processada: " + tx.hash + "  - Registro salvo na Blockchain. Status: " + txReceipt.status);
    } else {
      alert("Transação processada: " + tx.hash + "  - Mas houve um erro na blockchain. Veja pelo etherscan");
    }
  } catch (err) {
    console.error(err);
    alert("Houve um erro ao salvar o registro do contrato de aluguel");
  }
}