package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Opentimer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.814 38.195a16.153 16.153 0 0 1-16.153 0m8.077-3.275v-1.714m-5.352.28L19.243 32m-4.775-2.432l1.485-.857m-2.919-4.495h1.714m-.28-5.351l1.485.857m2.433-4.775l.857 1.484m4.495-2.918v1.714m5.352-.28l-.857 1.484m4.774 2.434l-1.484.857m2.918 4.494h-1.714m.28 5.352l-1.484-.857m-2.433 4.775L28.233 32"/><circle cx="23.738" cy="24.206" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.591 24.06l-5.356-5.356m-6.501 16.31a16.153 16.153 0 1 1 24.378-.425m-24.378.425c-1.842 1.906-3.747 3.657-2.91 8.486h7.056a11.902 11.902 0 0 1-.22-5.305"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.814 38.195c1.044 1.06 1.677 2.642 1.446 5.305h7.14c.202-3.85-1.548-6.603-4.288-8.911M32.704 10.75c1.97.976 2.92-2.5.866-3.536c-7.235-3.648-12.367-3.59-19.743 0c-2.144 1.043-.995 4.39.897 3.59"/>`),
		g.Group(children),
	)
}