package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.198 21.577s-.547 7.4.027 10.29c.497 2.495 4.625 4.214 4.625 4.214v3.83c-1.34.138-5.5 1.807-7.033 4.11h18.426c-1.532-2.303-5.692-3.973-7.032-4.11v-3.83s7.57.136 8.665-4.215a40.327 40.327 0 0 0 .592-7.274"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.85 38.133a56.78 56.78 0 0 1-6.403-.547c-1.774-.341-8.538-.22-8.538-15.38c0-15.817 10.453-17.185 17.459-17.185c11 0 16.475 5.472 17.148 12.621c.462 4.905-2.7 14.361-4.177 18.028a5.524 5.524 0 0 1-1.368-1.9c-.858.274-1.697 4.144-1.697 4.144l-5.637.44l-.81-2.358"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.02 13.23a14.996 14.996 0 0 0-2.683 8.648s3.01.493 8.429-5.09a4.497 4.497 0 0 1 1.587.772a7.98 7.98 0 0 0 .858 5.96s6.038-6.513 6.585-10.016"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.926 20.729c-1.04.547-4.597 5.363-4.597 5.363a17.355 17.355 0 0 1 0-4.761s-.75-3.362-3.03-4.39"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.329 21.33c.164-2.08 1.97-4.05 2.517-5.199m-13.08.657A20.464 20.464 0 0 0 25.21 12.3m2.112 21.578c-1.108 0-2.791-.247-3.53-1.437M20.85 21.516c2.408 0 3.64.91 3.64 4.316s-1.56 4.023-3.037 4.023s-2.71-.868-2.71-4.17c0-2.726.8-4.169 2.107-4.169Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.56 21.85c-1.272.616-1.436 2.334-1.436 3.836s-.156 3.026 1.03 4.11m11.146-8.28c-.655.006-2.349.937-2.267 4.576s1.667 3.741 2.132 3.704a1.145 1.145 0 0 0 1.125-.75"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.329 26.092a6.481 6.481 0 0 1-.67 3.709m-21.332-13.67A47.108 47.108 0 0 0 12.04 28.5a16 16 0 0 0 3.612 6.24a8.88 8.88 0 0 1-9.166-5.807m31.44-18.134a3.867 3.867 0 0 0-3.458.954M20.85 38.706a6.714 6.714 0 0 0 2.213.616a5.504 5.504 0 0 0 2.147-.462M33.164 24a4.94 4.94 0 0 0-.675-2.169"/>`),
		g.Group(children),
	)
}