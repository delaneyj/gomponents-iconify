package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PublicServices(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.072 34.313a9.686 9.686 0 0 0 3.132 2.98c3.094 1.818 6.176 3.639 9.237 5.444a8.97 8.97 0 0 0 9.178 0c3.06-1.805 6.142-3.627 9.236-5.445c2.879-1.697 4.698-4.75 4.721-8.02c.026-3.515.026-7.03 0-10.544c-.023-3.27-1.842-6.324-4.721-8.02c-2.322-1.365-6.945-4.094-9.236-5.445a8.97 8.97 0 0 0-9.178 0c-3.061 1.805-6.143 3.626-9.237 5.445a9.704 9.704 0 0 0-3.048 2.856"/><rect width="3.888" height="5.151" x="9.797" y="16.272" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.944" ry="1.944"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.807 20.444a1.943 1.943 0 0 1-1.688.98h0a1.944 1.944 0 0 1-1.944-1.945v-1.263c0-1.074.87-1.944 1.944-1.944h0c.72 0 1.35.393 1.686.976M14.89 28.172a1.943 1.943 0 0 1-1.687.98h0a1.944 1.944 0 0 1-1.944-1.945v-1.263c0-1.074.87-1.944 1.944-1.944h0c.72 0 1.35.393 1.686.976m-9.485-3.553v-5.151h3.033m12.186 12.879V24H17.59c0 3.303-.614 5.06-.614 5.06m11.771.091V24h3.032M7.916 29.332L5.772 24m4.025 0l-2.555 7.243m17.723-1.911L22.822 24m4.024 0l-2.555 7.243M33.549 24v3.207c0 1.074.87 1.944 1.943 1.944h0c1.074 0 1.944-.87 1.944-1.944V24m0 3.207v1.944"/>`),
		g.Group(children),
	)
}