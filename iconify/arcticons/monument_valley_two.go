package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonumentValleyTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.133 16.648a4.142 4.142 0 0 0 2.997 6.83"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.981 16.646a3.139 3.139 0 0 0-1.235.208c-3.201 1.365-6.967 3.609-4.308 10.491c0 0 .663-3.074 2.656-3.588l.002-.005a4.14 4.14 0 1 0 2.885-7.106Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.5 24.9c-.33 3.136-3.985 7.704-3.985 7.704s-.04 3.553 4.69 3.553s4.405-2.944 4.405-3.411s-2.748-6.379-2.856-8.214"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.065 42.471l-1.14-.913l.139-5.479m3.617 4.951l-1.319-.995l.077-3.956"/><circle cx="22.133" cy="29.112" r="1.259" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.13 28.917a8.514 8.514 0 0 0 5.915.827m-6.417-2.137c-1.266-.383-5.363 6.273-6.421 6.944"/><circle cx="30.26" cy="12.118" r="4.501" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.98 16.61a4.501 4.501 0 0 0-2.817-7.747"/><circle cx="36.83" cy="7.301" r="2.801" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m33.821 9.368l.708-.474m-4.113 16.171c-.027 1.38 1.625 10.882 7.31 10.232a5.207 5.207 0 0 0 0-5.864a5.199 5.199 0 0 0 4.737.516c-.27-1.149-2.978-3.774-2.978-3.774a3.351 3.351 0 0 0 2.978.135c-1.272-.65-9.719-4.846-10.206-7.011a1.53 1.53 0 0 1-2.49 0c-.054.785-2.003 13.02-3.249 14.618c1.597.406 2.304.541 2.75-.135s1.23-5.712 1.148-8.717Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.332 41.957l1.928-1.191l-.56-8.465h2.944M31.499 43.5l1.868-1.543l-.164-8.839m-9.983-4.639c1.322-1.032 4.337-2.388 6.187-6.759m2.85-2.421V16.15m-2.49 3.149v-2.651"/>`),
		g.Group(children),
	)
}