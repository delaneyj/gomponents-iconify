package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tuxrider(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.756 40.511q-.357.022-.714.022a11.249 11.249 0 0 1-1.128-.058m-9.538-5.261a10.87 10.87 0 0 1-2.782-6.816c0-4.737 2.395-9.315 2.395-12.348S14.799 4.5 24.326 4.5c9.314 0 10.766 9.676 10.964 12.614c.32 4.737 2.182 7.079 2.182 10.964a11.437 11.437 0 0 1-2.762 7.257"/><circle cx="24.033" cy="29.495" r="3.925" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.255 17.374a93.483 93.483 0 0 0-10.923 8.209m-11.465-8.209c3.502 1.705 9.303 5.655 11.166 8.196m.945-8.099a3.474 3.474 0 0 1-.446-1.774a3.29 3.29 0 0 1 3.426-3.632a3.964 3.964 0 0 1 3.68 3.852a3.667 3.667 0 0 1-3.234 3.413a3.677 3.677 0 0 1-.672-.072"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.308 19.063a3.485 3.485 0 0 1-3.48-3.865c0-3.094 2.122-4.59 3.852-4.59c2.089 0 3.852 2.395 3.852 5.089a3.317 3.317 0 0 1-.402 1.664m7.148 17.497c4.112 0 6.786 1.616 6.786 4.55c0 1.876-3.453 1.876-6.786 1.876s-6.626.059-6.626-1.915c0-2.495 3.033-4.511 6.626-4.511Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.848 40.122a1.271 1.271 0 0 1 .615 1.462c-.15.64-2.714 1.916-7.504 1.916c-4.504 0-6.256-.388-6.203-2.989q.005-.248.031-.523m-8.394-5.13c-4.112 0-6.786 1.616-6.786 4.55c0 1.876 3.453 1.876 6.786 1.876s6.626.059 6.626-1.915c0-2.495-3.034-4.511-6.626-4.511Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.823 40.122a1.26 1.26 0 0 0-.615 1.462c.17.682 2.714 1.916 7.504 1.916c4.525 0 6.271-.392 6.202-3.025q-.006-.231-.03-.487m13.06-18.991c.145.533 7.689 4.846 6.77 7.294c-.744 1.985-4.056-2.077-5.433-2.361m-25.322-4.933c-.171.897-7.59 4.846-6.672 7.294c.744 1.985 4.056-2.077 5.434-2.361m13.611-8.556c1.537 0 3.4 1.884 3.4 1.884l-3.4 4.163S20.14 19.81 20.18 19.33s1.915-1.956 4.15-1.956Zm-10.956 17.84c-.76-2.691-.56-13.124 5.787-16.15m14.692 16.044c.836-1.667 2.476-11.261-4.431-16.045"/><circle cx="26.293" cy="15.585" r=".75" fill="currentColor"/><circle cx="22.981" cy="15.697" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.033 30.757v-2.523"/>`),
		g.Group(children),
	)
}