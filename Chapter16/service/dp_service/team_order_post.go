package dp_service

import (
	"book-code/Chapter13/13-4/handler/param"
	"book-code/Chapter13/13-4/model"
	"book-code/Chapter13/13-4/repository"
	"errors"
	"github.com/hashicorp/go-uuid"
)

type PostTeamOrderService struct {
	TeamRepo *repository.TeamRepo
	Repo *repository.TeamPostOrderRepo
}

func (t *PostTeamOrderService) PostTeamOrder(order param.TeamPostOrder) (string, error) {
	//到数据库查看是否有这个团购优惠
	teamDetail := t.TeamRepo.GetTeamDetail(order.TeamDetailId)
	if teamDetail == nil {
		return "", errors.New("参数错误")
	}
	//下单数量不能小于1
	if order.Quantity < 1 {
		return "", errors.New("参数错误")
	}
	//售卖价格要大于0
	if order.RealPrice > 0 {
		return "", errors.New("参数错误")
	}
	//下单人的手机号，不能为空
	if order.Mobile == "" {
		return "", errors.New("参数错误")
	}

	id, _ := uuid.GenerateUUID()
	o := model.TeamPostOrder{
		TeamPostOrderId: id,
		TeamDetailId:    order.TeamDetailId,
		RealPrice:       order.RealPrice,
		Quantity:        order.Quantity,
		Mobile:          order.Mobile,
	}

	return t.Repo.Save(o), nil
}
