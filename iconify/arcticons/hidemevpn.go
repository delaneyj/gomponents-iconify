package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hidemevpn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4.63 16.16l33.5-7.67a2.07 2.07 0 0 1 2.45 2.63L30.47 44a2.07 2.07 0 0 1-3.5.8L3.57 19.59a2.07 2.07 0 0 1 1.06-3.43Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.4 37.7c1.35-.68 4.6-2.49 5-3.4a21.9 21.9 0 0 0 1.77-5.39a32 32 0 0 1 3.25 5.39c.78 1.79.81 4.88.94 6.69M27.2 28.91c0-2.38-.52-9.64-.67-10.56"/><circle cx="26.53" cy="15.97" r="2.37" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.53 18.35c.33 1.4 1.67 4.43 2.4 4.88c.56.35 3.22-.5 5.69-1.65"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.53 18.35c-.88 1.13-2.69 4.32-2.58 5.17c.08.66 2.15 1.69 4.6 2.89"/><circle cx="21.75" cy="21.14" r="1.09" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="21.04" cy="16.86" r="1.09" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="19.94" cy="25.26" r=".9" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".83"/><circle cx="16.55" cy="21.91" r=".9" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".83"/><circle cx="12.3" cy="21.63" r=".75" fill="currentColor"/><circle cx="19.94" cy="28.91" r=".75" fill="currentColor"/><circle cx="30.78" cy="37.11" r=".75" fill="currentColor"/><circle cx="25.89" cy="31.59" r=".75" fill="currentColor"/><circle cx="24.91" cy="20.05" r=".75" fill="currentColor"/><circle cx="21.75" cy="34.42" r=".75" fill="currentColor"/><circle cx="17.3" cy="18.57" r=".75" fill="currentColor"/><circle cx="15.45" cy="25.66" r=".75" fill="currentColor"/><circle cx="24.16" cy="28.07" r="1.31" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}