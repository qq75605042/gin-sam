package controllers

import "gin-sam/services"

var (
	userFService       services.UserService
	blackFService      services.BlackService
	blackIpFService    services.BlackIpService
	couponFService     services.CouponService
	giftFService       services.GiftService
	userDayNumFService services.UserDayNumService
)

func Setup() {
	userFService = services.NewUserService()
	blackIpFService = services.NewBlackIpService()
	blackFService = services.NewBlackService()
	couponFService = services.NewCouponService()
	giftFService = services.NewGiftService()
	userDayNumFService = services.NewUserDayNumService()
}
