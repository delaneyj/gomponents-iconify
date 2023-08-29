package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KissingFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#ffe500" d="M1.5 24.83a15.33 15.33 0 1 0 30.66 0a15.33 15.33 0 1 0-30.66 0Z"/><path fill="#ebcb00" d="M16.83 9.5a15.33 15.33 0 1 0 15.33 15.33A15.33 15.33 0 0 0 16.83 9.5Zm0 28.36a14 14 0 1 1 14-14a14 14 0 0 1-14 14Z"/><path fill="#fff48c" d="M12.23 12.57a4.6 1.15 0 1 0 9.2 0a4.6 1.15 0 1 0-9.2 0Z"/><path fill="#45413c" d="M4.57 45.5a12.26 1.5 0 1 0 24.52 0a12.26 1.5 0 1 0-24.52 0Z" opacity=".15"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M1.5 24.83a15.33 15.33 0 1 0 30.66 0a15.33 15.33 0 1 0-30.66 0Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M16.83 29.63a1.53 1.53 0 1 1 .77 2.86a1.54 1.54 0 0 1 0 3.07a1.57 1.57 0 0 1-.78-.21"/><path fill="#ffaa54" d="M5.72 28.66c0 .64.85 1.15 1.91 1.15s1.92-.51 1.92-1.15s-.86-1.15-1.92-1.15s-1.91.49-1.91 1.15Zm22.22 0c0 .64-.86 1.15-1.91 1.15s-1.92-.51-1.92-1.15s.89-1.15 1.89-1.15s1.94.49 1.94 1.15Z"/><path fill="#45413c" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M9.55 24.45a.77.77 0 1 1-.77-.77a.77.77 0 0 1 .77.77Zm14.56 0a.77.77 0 1 0 .77-.77a.76.76 0 0 0-.77.77Z"/><path fill="none" stroke="#48cf3e" stroke-linecap="round" stroke-linejoin="round" d="M37.88 17a1.65 1.65 0 1 1-1.64-1.64A1.65 1.65 0 0 1 37.88 17Zm8.21-1.69a1.65 1.65 0 1 1-1.64-1.65a1.65 1.65 0 0 1 1.64 1.65Z"/><path fill="none" stroke="#48cf3e" stroke-linecap="round" stroke-linejoin="round" d="M37.88 16.95V9.56l8.21-3.29v9.03"/><path fill="none" stroke="#ff4064" stroke-linecap="round" stroke-linejoin="round" d="M36.11 40.55a1.62 1.62 0 1 1-1.61-1.61a1.62 1.62 0 0 1 1.61 1.61Zm0 0v-7.27l2.42 1.62"/><path fill="none" stroke="#ffb700" stroke-linecap="round" stroke-linejoin="round" d="M27.1 7.36a1 1 0 1 1-1.05-1a1 1 0 0 1 1.05 1Z"/><path fill="none" stroke="#ffb700" stroke-linecap="round" stroke-linejoin="round" d="M27.1 7.36v-4.7l1.57 1.04"/><path fill="none" stroke="#00aed9" stroke-linecap="round" stroke-linejoin="round" d="M38.81 23.83v-.58m1.22 1.08l.4-.4m.1 1.62h.58m-1.08 1.22l.4.41m-1.62.1v.57m-1.22-1.08l-.41.41m-.1-1.63h-.57m1.08-1.22l-.41-.4M44.47 32v-3.06m-1.53 1.53h3.07m-4.9-26.28V2.66m-.77.77h1.53m.28 35.83v-1.54m-.77.77h1.53M31.23 9.37V7.83m-.77.77h1.53"/>`),
		g.Group(children),
	)
}