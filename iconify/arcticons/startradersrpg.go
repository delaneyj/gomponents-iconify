package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Startradersrpg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m41.15 24.77l2.35.23l-2.12-11.17l-10.65-1.53l-11.63 2.17l-3.53 11.39L27 30.22l3.73-17.92"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.94 34.59l10.79-5.14l.53-5.91l-12.81 3.22Zm0 0l-3.53-2M28.88 34l-.52 2.1c-.76.4-1.27 0-1.73-.69l-.44-5.48"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.63 35.41c-.64.3-1.49 1-2.1-.78l-.42-5.48"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.53 34.63c-1.44.77-2.09.32-2.1-1.13l-.23-5.1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.36 32L9.17 24.35l1.88-1.28l4.83 1.83m1.22-3.98l-5.47-2.62l-4-.08s-.45-.54 0-.54l4.23-.07l6.27.07M35 20l-.17 3.49L36 24.65l1-.25l.93-1.28l-.27-3.34Zm-5.06-3.95l12 .46m-27.53 1.12c-.1-1.22-.11-2.9 1.32-3.44c1.17-.73 2.26 0 3.37.28"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m11.28 17.62l.3-4l4.1-.28l.05.89m-1.31 13.18l-.53 5.8l-2.19-7.39m-2.53-1.47L4.5 25.56L9.32 22l1.73 1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m11.45 15.38l-.83-1C8.88 14.42 9.27 16 9.29 17l2.16-1.57"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8 17.67c-.43-2 .24-3 1.67-3m-2.16 8.7A2.73 2.73 0 0 1 6 20.51m1.65.66l-3.13-1.31l3.59-.21Zm-1.36-1.42C5.74 18.59 6.23 18 7.45 18m.2 3.17l1.67.83m-1.21-2.35l1.68-1.38"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.34 15.13a3.77 3.77 0 0 0-1.44-3.06a9.5 9.5 0 0 1 3.31 2.39m5.47-1.16l-.28-1.51h2.39l3 2.43M31.66 23l-.12-.88M40.19 21l.15.72"/>`),
		g.Group(children),
	)
}