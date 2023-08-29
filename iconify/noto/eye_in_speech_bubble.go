package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeInSpeechBubble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<defs><path id="notoEyeInSpeechBubble0" d="M64.73 20C19.51 20 5 61.27 5 61.27s12.42 39.85 59.73 39.85c7.34.06 14.64-1.02 21.64-3.21c6 8.49 18.78 12.47 25.3 13.33h.27c1.1-.02 1.99-.93 1.97-2.03a2.04 2.04 0 0 0-.33-1.07c-2.93-4.31-7.58-12-9.22-19.67a60.533 60.533 0 0 0 19.78-27.21S109.95 20 64.73 20z" opacity=".2"/></defs><g fill="#FFF"><use href="#notoEyeInSpeechBubble0" opacity=".2"/><path d="M104.36 88.48a60.533 60.533 0 0 0 19.78-27.21S109.95 20 64.73 20S5 61.27 5 61.27s12.42 39.85 59.73 39.85c7.34.06 14.64-1.02 21.64-3.21c6 8.49 18.78 12.47 25.3 13.33c1.1.13 2.09-.65 2.23-1.74c.06-.47-.05-.95-.32-1.35c-2.93-4.31-7.58-12.03-9.22-19.67z"/></g><circle cx="64.72" cy="60.54" r="28.76" fill="#3F474C"/><circle cx="64.72" cy="60.54" r="16.01" fill="#231F20"/><circle cx="81.92" cy="56.34" r="6.8" fill="#F5F8FA"/><g fill="none"><use href="#notoEyeInSpeechBubble0" opacity=".2"/><path stroke="#000" stroke-linejoin="round" stroke-width="6" d="M104.36 88.48a60.533 60.533 0 0 0 19.78-27.21S109.95 20 64.73 20S5 61.27 5 61.27s12.42 39.85 59.73 39.85c7.34.06 14.64-1.02 21.64-3.21c6 8.49 18.78 12.47 25.3 13.33c1.1.13 2.09-.65 2.23-1.74c.06-.47-.05-.95-.32-1.35c-2.93-4.31-7.58-12.03-9.22-19.67z"/></g>`),
		g.Group(children),
	)
}