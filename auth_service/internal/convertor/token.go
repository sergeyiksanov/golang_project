package convertor

import (
	"github.com/sergeyiksanov/golang_project/auth_service/internal/dto"
	"github.com/sergeyiksanov/golang_project/auth_service/internal/entity"
)

func TokenEntityToTokenDto(t entity.Token) dto.TokenDto {
	return dto.TokenDto{
		JTI:       t.JTI,
		SubjectId: t.SubjectId,
		TokenType: t.TokenType,
		Revoked:   t.Revoked,
		IssuedAt:  t.IssuedAt,
		ExpiresAt: t.ExpiresAt,
	}
}

func TokenDtoToTokenEntity(t dto.TokenDto) entity.Token {
	return entity.Token{
		JTI:       t.JTI,
		SubjectId: t.SubjectId,
		TokenType: t.TokenType,
		Revoked:   t.Revoked,
		IssuedAt:  t.IssuedAt,
		ExpiresAt: t.ExpiresAt,
	}
}
