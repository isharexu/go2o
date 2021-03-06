/**
 * Copyright 2014 @ Ops Inc.
 * name :
 * author : newmin
 * date : 2013-12-03 23:20
 * description :
 * history :
 */

package dps

import (
	"github.com/atnet/gof/app"
	"go2o/core/query"
	"go2o/core/repository"
)

var (
	Context         app.Context
	PromService     *promotionService
	ShoppingService *shoppingService
	MemberService   *memberService
	PartnerService  *partnerService
	SaleService     *saleService
)

func Init(ctx app.Context) {
	Context := ctx
	db := Context.Db()

	/** Repository **/
	userRep := repository.NewUserRep(db)
	partnerRep := repository.NewPartnerRep(db, userRep)
	memberRep := repository.NewMemberRep(db)
	promRep := repository.NewPromotionRep(db, memberRep)
	saleRep := repository.NewSaleRep(db)
	deliveryRep := repository.NewDeliverRep(db)
	spRep := repository.NewShoppingRep(db, partnerRep, saleRep, promRep, memberRep,deliveryRep)

	/** Query **/
	mq := query.NewMemberQuery(db)
	pq := query.NewPartnerQuery(db)

	/** Service **/
	PromService = NewPromotionService(promRep)
	ShoppingService = NewShoppingService(spRep)
	MemberService = NewMemberService(memberRep, mq)
	PartnerService = NewPartnerService(partnerRep, pq)
	SaleService = NewSaleService(saleRep)
}
