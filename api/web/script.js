const chatBox = document.getElementById("chat-box");
const chatForm = document.getElementById("chat-form");
const userInput = document.getElementById("user-input");

const BOT_NAME = "Porto";

const CHAT_API_PATH = "/chat";

async function requestChatbotResponse(question) {
  return fetch(CHAT_API_PATH, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ question }),
  })
    .then((response) => response.json())
    .catch((error) => console.error("Error:", error));
}

function appendMessage(sender, message) {
  const msgDiv = document.createElement("div");
  msgDiv.className = "chat-message";
  msgDiv.innerHTML = `<strong>${sender}:</strong> ${message}`;
  chatBox.appendChild(msgDiv);
  chatBox.scrollTop = chatBox.scrollHeight;
}

chatForm.addEventListener("submit", async function (e) {
  e.preventDefault();
  const question = userInput.value.trim();
  if (!question) return;
  appendMessage("Kamu", question);
  userInput.value = "";

  request = await requestChatbotResponse(question);

  console.log(request.answer);

  appendMessage(BOT_NAME, request.answer);
});
