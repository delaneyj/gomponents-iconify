package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecyclingSymbol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M7 45.5a17 1.5 0 1 0 34 0a17 1.5 0 1 0-34 0Z" opacity=".15"/><path fill="#46b000" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M26.83 2.42h-6.19A5.41 5.41 0 0 0 16 5.12L9.83 15.78l9.66 5.55l9.2-15.65a2.16 2.16 0 0 0-1.86-3.26Z"/><path fill="#6dd627" d="m35.06 10.13l-2.87-5a5.41 5.41 0 0 0-4.68-2.7h-6.87A5.4 5.4 0 0 0 16 5v.12a1.88 1.88 0 0 1 3.23 0l6.22 10.55l-1.71 1l10.6 3.24l2.5-10.8Z"/><path fill="#9ceb60" d="m19.19 5.14l.64 1.08a5.41 5.41 0 0 1 2.65-.7h6.86a5.41 5.41 0 0 1 4.28 2.1l-1.43-2.5a5.41 5.41 0 0 0-4.68-2.7h-6.87A5.4 5.4 0 0 0 16 5v.12a1.88 1.88 0 0 1 3.19.02Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="m35.06 10.13l-2.87-5a5.41 5.41 0 0 0-4.68-2.7h-6.87A5.4 5.4 0 0 0 16 5v.12a1.88 1.88 0 0 1 3.23 0l6.22 10.55l-1.71 1l10.6 3.24l2.5-10.8Z"/><path fill="#46b000" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="m2.29 35.48l3.1 5.36a5.4 5.4 0 0 0 4.68 2.7H22V32.4l-17.81-.16a2.16 2.16 0 0 0-1.9 3.24Z"/><path fill="#6dd627" d="M4.85 24.51L2 29.49a5.39 5.39 0 0 0 0 5.41l3.43 5.94a5.38 5.38 0 0 0 4.57 2.7h.1a1.88 1.88 0 0 1-1.61-2.8l6-10.67l1.71 1l-2.5-10.8L3.11 23.5Z"/><path fill="#9ceb60" d="m7.53 25.79l7-2.12l-.79-3.41L3.11 23.5l1.74 1L2 29.49a5.39 5.39 0 0 0 0 5.41l1 1.79a5.45 5.45 0 0 1 .55-4Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M4.85 24.51L2 29.49a5.39 5.39 0 0 0 0 5.41l3.43 5.94a5.38 5.38 0 0 0 4.57 2.7h.1a1.88 1.88 0 0 1-1.61-2.8l6-10.67l1.71 1l-2.5-10.8L3.11 23.5Z"/><path fill="#46b000" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="m43.16 40.22l3.09-5.36a5.39 5.39 0 0 0 0-5.41L40 18.55l-9.63 5.6L39.4 40.2a2.17 2.17 0 0 0 3.76.02Z"/><path fill="#6dd627" d="M32.37 43.49h5.77a5.4 5.4 0 0 0 4.68-2.7l3.43-5.94a5.4 5.4 0 0 0 0-5.32v-.09a1.88 1.88 0 0 1-1.62 2.79l-12.26.11v-2l-8.1 7.56l8.1 7.57Z"/><path fill="#9ceb60" d="M45.84 35.47a5.43 5.43 0 0 0 .46-5.93v-.09a1.88 1.88 0 0 1-1.62 2.79l-12.26.11v-2l-8.1 7.56L26 39.55l4.23-3.94Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M32.37 43.49h5.77a5.4 5.4 0 0 0 4.68-2.7l3.43-5.94a5.4 5.4 0 0 0 0-5.32v-.09a1.88 1.88 0 0 1-1.62 2.79l-12.26.11v-2l-8.1 7.56l8.1 7.57Z"/>`),
		g.Group(children),
	)
}