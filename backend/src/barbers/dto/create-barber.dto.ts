import { IsString, IsNotEmpty, IsUrl } from 'class-validator';
import { ApiProperty } from '@nestjs/swagger';

export class CreateBarberDto {
  @ApiProperty({ description: 'Name of the barber' })
  @IsString()
  @IsNotEmpty()
  readonly name: string;

  @ApiProperty({ description: 'Photo url of the barber' })
  @IsUrl()
  @IsNotEmpty()
  readonly photo: string;
}
