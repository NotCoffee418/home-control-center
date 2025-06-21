// components/Nav.tsx
import { Group, Button, Container } from '@mantine/core';
import { Link, useLocation } from 'react-router-dom';

const TopNav = () => {
  const { pathname } = useLocation();
  
  const TopNavItem = ({ to, children }: { to: string; children: string }) => (
    <Button
      component={Link}
      to={to}
      variant={pathname === to || (to === '/dashboard' && pathname === '/') ? 'filled' : 'subtle'}
    >
      {children}
    </Button>
  );

  return (
    <Container size="xl" py="md">
      <Group>
        <TopNavItem to="/dashboard">Dashboard</TopNavItem>
        <TopNavItem to="/profiles">Profiles</TopNavItem>
      </Group>
    </Container>
  );
};

export default TopNav;