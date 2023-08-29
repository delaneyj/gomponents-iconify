package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Directa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.042 42.54c-4.994.23-9.988.458-14.982.68a47.37 47.37 0 0 0-.097-4.28c-2.295 3.627-5.613 5.26-9.826 4.28c-4.998-1.16-7.989-4.627-8.854-9.728c-.57-3.361-.44-6.749.779-9.924c2.57-6.696 7.623-9.262 14.012-6.127c1.12.55 2.48 1.474 3.305 2.43c.097-4.345.13-10.336.097-14.682" opacity=".998"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.196 5.278c4.97-.035 9.9-.294 14.787-.778h2.335v38.165M22.085 21.587c1.236-.29 2.098.647 2.312 1.911c.65 3.841.75 7.772.157 11.63c-.122.794-.34 1.587-.796 2.23c-.599.842-1.629.968-2.23.16c-.397-.533-.44-1.247-.479-1.912c-.24-4.137-.268-8.293 0-12.427c.048-.745.363-1.434 1.036-1.592Z" opacity=".998"/>`),
		g.Group(children),
	)
}