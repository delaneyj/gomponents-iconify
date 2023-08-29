package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mirakel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.355 17.943a4.203 4.203 0 0 0 6.87-.234"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.937 36.25c-.94-1.882-4.416-1.036-4.416 1.985c0 3.399 3.196 4.394 5.341 4.394s5.636-1.022 5.636-5.06s-2.239-3.659-2.239-6.483s3.372-4.698 3.372-7.984s-4.371-5.393-4.371-10.892a6.574 6.574 0 0 1 6.94-6.839c5.467 0 8.894 4.82 8.894 9.363s-2.996 4.95-2.996 7.96a4.396 4.396 0 0 0 1.013 2.705c-1.97 1.846-4.052 2.93-4.052 6.544c0 2.27.379 6.565 5.678 6.565c5.226 0 7.6-4.734 8.064-7.426s2.38-5.556 5.1-5.556c3.28 0 3.89 2.338 3.89 3.973c0 2.63-3.488 4.256-5.797 1.062"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.287 36.21a2.357 2.357 0 0 1 2.44-2.742c1.822 0 2.705 1.662 2.705 3.302c0 1.284-.729 4.5-7.251 4.5c-4.17 0-5.758-3.216-7.252-5.879m9.695-20.504a2.382 2.382 0 0 0-3.661 1.949c0 1.654.717 4.822 4.768 3.73m-21.349 4.027l-5.247-5.247l-5.306 5.307l7.588 7.588"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.177 20.305c-1.13-1.644-2.435-.619-2.918.078c-.535.771.164 3.9 2.631 4.21m16.221.805l11.457-11.457l5.406 5.405l-17.12 17.12M8.521 27.345c-.593 2.213-.788 7.086 6.017 6.752m2.889 2.153l4.321 4.322l1.404-1.404"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.829 28.22c-4.054 2.006-3.18 7.836 1.09 9.287a13.144 13.144 0 0 0 8.567-.44"/><circle cx="16.763" cy="13.312" r="2.254" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="23.335" cy="13.312" r="2.254" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="23.335" cy="13.312" r=".75" fill="currentColor"/><circle cx="16.763" cy="13.312" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.096 36.38c-1.71-2.238-3.14-3.597-5.103-3.044c-1.886.532-2.621 1.766-2.495 4.233"/>`),
		g.Group(children),
	)
}