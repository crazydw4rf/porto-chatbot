const chatBox = document.getElementById("chat-box");
const chatForm = document.getElementById("chat-form");
const userInput = document.getElementById("user-input");

const BOT_NAME = "Porto";

const CHAT_API_PATH = "/chat";

async function requestChatbotResponse(prompt) {
  return fetch(CHAT_API_PATH, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ prompt }),
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

/**
 * @typedef {Object} ChatResponse
 * @property {string} response - The response from the chatbot.
 * @property {string} error - Any error message if the request fails.
 */

chatForm.addEventListener("submit", async function (e) {
  e.preventDefault();
  const question = userInput.value.trim();
  if (!question) return;
  appendMessage("Kamu", question);
  userInput.value = "";

  /** @type {ChatResponse} */
  const request = await requestChatbotResponse(question);

  if (request.error) {
    appendMessage(BOT_NAME, `Error: ${request.error}`);
    return;
  }

  appendMessage(BOT_NAME, request.response);
});
