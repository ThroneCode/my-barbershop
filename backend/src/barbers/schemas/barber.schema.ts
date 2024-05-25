import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Document } from 'mongoose';
import { ApiProperty } from '@nestjs/swagger';

@Schema()
export class Barber extends Document {
  @ApiProperty({ description: 'ID of the barber' })
  _id: string;

  @Prop({ required: true })
  @ApiProperty({ description: 'Name of the barber' })
  name: string;

  @Prop()
  @ApiProperty({ description: 'Photo url of the barber' })
  photo: string;

  @Prop({ default: false })
  @ApiProperty({ description: 'Is the barber working today?' })
  is_working: boolean;
}

export const BarberSchema = SchemaFactory.createForClass(Barber);
