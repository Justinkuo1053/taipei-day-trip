// =====================
// model: 取得訂單資料
// =====================
// 根據網址上的訂單編號，向後端取得訂單資訊
async function fetchOrderData(orderNumber) {
  const response = await fetch("/api/order/" + orderNumber, {
    method: "GET",
    headers: {
      // 需帶上登入者 token，後端才能驗證
      Authorization: "Bearer " + localStorage.getItem("user_token"),
    },
  });
  const data = await response.json();
  console.log("fetchOrderData response:", data);
  // 若查無訂單或未登入，導回首頁
  if (data.error) {
    location.href = "/";
  } else {
    // 成功取得訂單，渲染訊息與 footer
    console.log(data);
    renderMessage(data.data.status, orderNumber);
    renderFooter();
  }
}

// =====================
// view: 畫面渲染
// =====================
// 動態調整 footer 位置，讓頁面底部美觀
function renderFooter() {
  const footer = document.querySelector("footer");
  footer.style.height =
    window.innerHeight - footer.getBoundingClientRect().top + "px";
}

// 根據訂單狀態顯示成功/失敗訊息與訂單編號
function renderMessage(status, orderNumber) {
  let result = document.querySelector(".order__result");
  let error = document.querySelector(".order__error");
  console.log("renderMessage called, status:", status, "orderNumber:", orderNumber);
  if (status == 0 || status == 1) {
    // 付款成功
    result.innerHTML = "行程預定成功！";
  } else {
    // 付款失敗
    result.innerHTML = "行程預定失敗！";
    error.style.display = "block";
    document.querySelector(".order__note").style.display = "none";
    document.querySelector(".order__number").style.display = "none";
    document.querySelector(".order__alert").style.display = "none";
  }
  // 顯示訂單編號
  document.querySelector(".order__number").innerHTML = orderNumber;
}

// =====================
// controller: 頁面初始化
// =====================
// 先調整 footer 位置
renderFooter();
// 從網址取得訂單編號（?number=xxxx）
const orderNumber = location.href.match(/^.+\?number=(.+)$/);
// 若有訂單編號才查詢訂單資料
if (orderNumber !== null) {
  fetchOrderData(orderNumber[1]);
}
