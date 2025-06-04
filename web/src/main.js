import './style.css'

const app = document.querySelector('#app');
let isLoading = false;
let shortenedUrl = '';

function render() {
  app.innerHTML = `
    <main class="min-h-screen flex items-center justify-center bg-gray-50 px-4">
      <div class="w-full max-w-md space-y-8">
        <div class="text-center">
          <h1 class="text-6xl font-bold text-gray-900 tracking-tight">morph</h1>
          <p class="mt-2 text-sm text-gray-600">simple URL shortening</p>
        </div>

        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <form id="shortenForm" class="space-y-4" ${isLoading ? 'aria-busy="true"' : ''}>
            <div>
              <label for="urlInput" class="sr-only">Enter URL to shorten</label>
              <input
                type="url"
                id="urlInput"
                placeholder="https://example.com/very-long-url"
                class="w-full px-3 py-3 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent text-sm"
                required
                ${isLoading ? 'disabled' : ''}
              />
            </div>

            <button
              type="submit"
              class="w-full bg-black text-white py-3 px-4 rounded-md hover:bg-violet-800 focus:outline-none focus:ring-2 focus:ring-black focus:ring-offset-2 transition-colors text-sm font-medium disabled:opacity-50 disabled:cursor-not-allowed"
              ${isLoading ? 'disabled' : ''}
            >
              ${isLoading ? 'Shortening...' : 'Shorten URL'}
            </button>
          </form>

          ${shortenedUrl ? `
            <div class="mt-6 pt-6 border-t border-gray-200">
              <div class="flex items-center justify-between bg-gray-50 rounded-md p-3">
                <span class="text-sm text-gray-900 font-mono truncate flex-1 mr-3">${shortenedUrl}</span>
                <button
                  id="copyBtn"
                  class="flex-shrink-0 text-violet-600 hover:text-violet-700 text-sm font-medium focus:outline-none focus:underline"
                >
                  Copy
                </button>
              </div>
            </div>
          ` : ''}
        </div>

        <div class="text-center text-xs text-gray-500">
          <p>free, fast, and simple URL shortening</p>
        </div>
      </div>
    </main>
  `;

  attachEventListeners();
}

function attachEventListeners() {
  const form = document.getElementById('shortenForm');
  const copyBtn = document.getElementById('copyBtn');

  if (form) {
    form.addEventListener('submit', handleSubmit);
  }

  if (copyBtn) {
    copyBtn.addEventListener('click', handleCopy);
  }
}

async function handleSubmit(e) {
  e.preventDefault();
  const urlInput = document.getElementById('urlInput');
  const url = urlInput.value.trim();

  if (!url) return;

  isLoading = true;
  render();

  try {
    // simulates an api call
    // TODO: implement api call
    await new Promise(resolve => setTimeout(resolve, 1000));

    const mockShortUrl = `https://morph.ly/${Math.random().toString(36).substr(2, 8)}`;

    shortenedUrl = mockShortUrl;
    isLoading = false;
    render();

  } catch (error) {
    isLoading = false;
    console.error('Error shortening URL:', error);
    render();
  }
}

async function handleCopy() {
  try {
    await navigator.clipboard.writeText(shortenedUrl);

    const copyBtn = document.getElementById('copyBtn');
    const originalText = copyBtn.textContent;
    copyBtn.textContent = 'Copied!';
    copyBtn.classList.add('text-green-600');

    setTimeout(() => {
      copyBtn.textContent = originalText;
      copyBtn.classList.remove('text-green-600');
    }, 2000);

  } catch (error) {
    console.error('Failed to copy:', error);
  }
}

render();
