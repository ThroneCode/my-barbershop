import {
  Controller,
  Post,
  Get,
  Body,
  UseGuards,
  UnauthorizedException,
  BadRequestException,
  Param,
  Delete,
  Put,
  HttpStatus,
} from '@nestjs/common';
import { AdminService } from './admin.service';
import { CreateAdminDto } from './dto/create-admin.dto';
import { Admin, AdminResponse } from './schemas/admin.schema';
import { JwtAuthGuard } from '../auth/jwt-auth.guard';
import { UpdateAdminDto } from './dto/update-admin.dto';
import { AuthService } from '../auth/auth.service';
import { GetJwtToken } from '../decorators/get-jwt-token.decorator';
import { ApiOperation, ApiProperty, ApiResponse } from '@nestjs/swagger';

@Controller('admin')
export class AdminController {
  constructor(
    private readonly adminService: AdminService,
    private readonly authService: AuthService,
  ) {}

  @Post()
  @ApiOperation({ summary: 'Create a new admin' })
  @ApiProperty({ type: CreateAdminDto, description: 'Admin details' })
  @ApiResponse({
    status: HttpStatus.CREATED,
    description: 'Return the created admin',
    type: Admin,
  })
  @ApiResponse({
    status: HttpStatus.UNAUTHORIZED,
    description: 'Invalid secret',
  })
  @ApiResponse({
    status: HttpStatus.BAD_REQUEST,
    description: 'Username already exists',
  })
  async create(@Body() createAdminDto: CreateAdminDto): Promise<AdminResponse> {
    if (createAdminDto.secret !== process.env.ADMIN_SECRET) {
      throw new UnauthorizedException('Invalid secret');
    }
    if (await this.adminService.findByUsername(createAdminDto.username)) {
      throw new BadRequestException('Username already exists');
    }
    return this.adminService.create(createAdminDto);
  }

  @Get()
  @ApiOperation({ summary: 'Get all admins' })
  @ApiResponse({
    status: HttpStatus.OK,
    description: 'Return all admins',
    type: [AdminResponse],
  })
  @UseGuards(JwtAuthGuard)
  async findAll(): Promise<AdminResponse[]> {
    return this.adminService.findAll();
  }

  @Get(':id')
  @ApiOperation({ summary: 'Get admin by id' })
  @ApiProperty({ type: String, description: 'Admin id' })
  @ApiResponse({
    status: HttpStatus.OK,
    description: 'Return the admin',
    type: AdminResponse,
  })
  @UseGuards(JwtAuthGuard)
  async findOne(@Param('id') id: string): Promise<AdminResponse> {
    return this.adminService.findOne(id);
  }

  @Delete(':id')
  @ApiOperation({ summary: 'Delete admin by id' })
  @ApiProperty({ type: String, description: 'Admin id' })
  @ApiResponse({
    status: HttpStatus.OK,
    description: 'Return the deleted admin',
    type: AdminResponse,
  })
  @ApiResponse({
    status: HttpStatus.BAD_REQUEST,
    description: 'Cannot delete own account',
  })
  @UseGuards(JwtAuthGuard)
  async remove(
    @Param('id') id: string,
    @GetJwtToken() token: string,
  ): Promise<AdminResponse> {
    const adminSession = await this.authService.getSessionData(token);
    if (adminSession._id === id) {
      throw new BadRequestException('Cannot delete own account');
    }
    return this.adminService.remove(id);
  }

  @Put(':id')
  @ApiOperation({ summary: 'Update admin by id' })
  @ApiProperty({ type: String, description: 'Admin id' })
  @ApiProperty({ type: UpdateAdminDto, description: 'Admin details' })
  @ApiResponse({
    status: HttpStatus.OK,
    description: 'Return the updated admin',
    type: Admin,
  })
  @UseGuards(JwtAuthGuard)
  async update(
    @Param('id') id: string,
    @Body() updateAdminDto: UpdateAdminDto,
  ): Promise<AdminResponse> {
    return this.adminService.update(id, updateAdminDto);
  }
}
