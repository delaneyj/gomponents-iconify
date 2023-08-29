package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StreamPiClient(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.737 16.614c-2.265.887-3.507 2.931-3.918 5.95c-3.324 3-4.075 5.856-.606 9.013c.188 5.083 3.578 8.078 9 9.75A5.499 5.499 0 0 0 24 43.5a5.499 5.499 0 0 0 4.787-2.173c5.422-1.671 8.812-4.667 9-9.75c3.469-3.157 2.718-6.014-.606-9.013c-.412-3.019-1.653-5.063-3.918-5.95"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.737 16.614c-3.688-1.835-4.929-6.33-5.651-10.794C12.772 4.782 24 1.404 24 12.254c0 3.541-4.16 6.899-9.263 4.36Zm18.526 0c3.689-1.835 4.929-6.33 5.651-10.794C35.228 4.782 24 1.404 24 12.254c0 3.541 4.16 6.899 9.263 4.36Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.405 12.615a17.332 17.332 0 0 0-7.89-4.614m13.281 4.614a17.332 17.332 0 0 1 7.889-4.614"/><ellipse cx="24.003" cy="28.954" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="9.511" ry="9.515"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.56 26.893v4.186a1.044 1.044 0 0 0 1.565.908l1.518-.873l.622-.356v-3.185a.619.619 0 0 0-.311-.537l-1.83-1.051a1.045 1.045 0 0 0-1.564.908Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.265 30.987v-4a1.15 1.15 0 0 1 1.724-.99l.02.012l3.467 1.989a1.143 1.143 0 0 1 .404 1.569a1.223 1.223 0 0 1-.404.409l-.02.011l-3.467 1.99a1.154 1.154 0 0 1-1.57-.419a1.137 1.137 0 0 1-.154-.571Z"/>`),
		g.Group(children),
	)
}