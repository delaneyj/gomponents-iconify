package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alligator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.58 20.27a23.487 23.487 0 0 1 8.103 9.714l3.495-.936a23.494 23.494 0 0 0-11.598-8.777Zm-.18-9.12l-1.052 3.924a23.478 23.478 0 0 1 14.41 5.595A23.49 23.49 0 0 0 22.4 11.15Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.477 5.982a1.72 1.72 0 0 0-1.767 1.275l-.327 1.22a1.716 1.716 0 0 0 .31 1.505l-.005.004a1.712 1.712 0 0 1 .298 1.481v.006l-.066.241l-.018.068l-1.175 4.383l-.034.13l-.272 1.012c-.943 2.123-2.076 2.97-4.077 2.752a4.705 4.705 0 0 0-4.53 5.979a4.275 4.275 0 0 1-.996 3.782L4.5 32.139v3.607a6.348 6.348 0 0 0 .123 1.236a6.207 6.207 0 0 0 .204.76c.02.06.043.12.065.18a6.24 6.24 0 0 0 .256.594c.018.037.034.076.053.113a6.256 6.256 0 0 0 .406.682l.012.016a6.258 6.258 0 0 0 .456.574l.114.127a6.263 6.263 0 0 0 .511.492l.04.035a6.26 6.26 0 0 0 .638.468l.084.05a6.231 6.231 0 0 0 .598.327c.045.022.089.045.135.065a6.228 6.228 0 0 0 .743.282a6.306 6.306 0 0 0 1.837.274h5.932l.894-.895c4.486-4.486 15.669-6.67 23.773-9.956a2.86 2.86 0 0 0 2.026-3.51a1.088 1.088 0 0 0-.349-.633l.08-.376l-2.764.741h-.006l-24.038 6.441l.351-1.31l1.065-3.97l3.544-13.223l1.683-6.283l.541-2.02l-2.02-.541l-.02-.004l-1.644-.441l-.076-.02l-.002.008a1.73 1.73 0 0 0-.268-.045Z"/><circle cx="12.342" cy="24.765" r="2.558" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="12.838" cy="24.343" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.38 9.563a1.005 1.005 0 0 1-.153-.527A1.076 1.076 0 0 1 20.33 7.99h0a1.159 1.159 0 0 1 .282.037"/>`),
		g.Group(children),
	)
}